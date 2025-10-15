import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as ec2 from 'aws-cdk-lib/aws-ec2';
import * as iam from 'aws-cdk-lib/aws-iam';

export class InfraStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const keyPairName = new cdk.CfnParameter(this, 'KeyPairName', {
      type: 'String',
      description: 'Name of an existing EC2 key pair to enable SSH access.',
    });

    const sshAllowedCidr = new cdk.CfnParameter(this, 'SshAllowedCidr', {
      type: 'String',
      default: '0.0.0.0/0',
      description: 'CIDR block allowed to access the instance via SSH (port 22).',
    });

    const appAllowedCidr = new cdk.CfnParameter(this, 'AppAllowedCidr', {
      type: 'String',
      default: '0.0.0.0/0',
      description: 'CIDR block allowed to reach the application port (8080).',
    });

    const vpc = new ec2.Vpc(this, 'AppVpc', {
      maxAzs: 2,
      natGateways: 0,
      subnetConfiguration: [
        {
          name: 'Public',
          subnetType: ec2.SubnetType.PUBLIC,
        },
      ],
    });

    const securityGroup = new ec2.SecurityGroup(this, 'AppSecurityGroup', {
      vpc,
      description: 'Security group for bookingapp EC2 instance',
      allowAllOutbound: true,
    });
    securityGroup.addIngressRule(
      ec2.Peer.ipv4(sshAllowedCidr.valueAsString),
      ec2.Port.tcp(22),
      'SSH access'
    );
    securityGroup.addIngressRule(
      ec2.Peer.ipv4(appAllowedCidr.valueAsString),
      ec2.Port.tcp(8080),
      'Application traffic'
    );

    const instanceRole = new iam.Role(this, 'AppInstanceRole', {
      assumedBy: new iam.ServicePrincipal('ec2.amazonaws.com'),
      description: 'IAM role for bookingapp EC2 instance',
    });
    instanceRole.addManagedPolicy(
      iam.ManagedPolicy.fromAwsManagedPolicyName('AmazonSSMManagedInstanceCore')
    );

    const userData = ec2.UserData.forLinux({ shebang: '#!/bin/bash' });
    userData.addCommands('set -euxo pipefail');
    userData.addCommands('dnf update -y');
    userData.addCommands('dnf install -y git docker tar');
    userData.addCommands('systemctl enable --now docker');
    userData.addCommands(
      'curl -SL https://github.com/docker/compose/releases/download/v2.29.2/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose'
    );
    userData.addCommands('chmod +x /usr/local/bin/docker-compose');
    userData.addCommands('curl -LO https://go.dev/dl/go1.24.0.linux-amd64.tar.gz');
    userData.addCommands('rm -rf /usr/local/go');
    userData.addCommands('tar -C /usr/local -xzf go1.24.0.linux-amd64.tar.gz');
    userData.addCommands('rm -f go1.24.0.linux-amd64.tar.gz');
    userData.addCommands('cat <<\'EOF\' >/etc/profile.d/bookingapp.sh');
    userData.addCommands('export PATH=$PATH:/usr/local/go/bin');
    userData.addCommands('EOF');
    userData.addCommands('id app >/dev/null 2>&1 || useradd -m app');
    userData.addCommands('mkdir -p /home/app/bookingapp');
    userData.addCommands('chown -R app:app /home/app');
    userData.addCommands('docker volume create bookingapp-mysql || true');

    const instance = new ec2.Instance(this, 'BookingAppInstance', {
      vpc,
      vpcSubnets: { subnetType: ec2.SubnetType.PUBLIC },
      instanceType: ec2.InstanceType.of(ec2.InstanceClass.T3, ec2.InstanceSize.SMALL),
      machineImage: ec2.MachineImage.latestAmazonLinux2023({
        edition: ec2.AmazonLinuxEdition.STANDARD,
      }),
      securityGroup,
      role: instanceRole,
      userData,
      keyName: keyPairName.valueAsString,
    });

    const elasticIp = new ec2.CfnEIP(this, 'BookingAppEip', {
      domain: 'vpc',
    });
    new ec2.CfnEIPAssociation(this, 'BookingAppEipAssociation', {
      eip: elasticIp.attrPublicIp,
      instanceId: instance.instanceId,
    });

    cdk.Tags.of(instance).add('Name', 'bookingapp-ec2');

    new cdk.CfnOutput(this, 'InstanceId', {
      value: instance.instanceId,
    });
    new cdk.CfnOutput(this, 'InstancePublicIp', {
      value: elasticIp.ref,
    });
    new cdk.CfnOutput(this, 'SecurityGroupId', {
      value: securityGroup.securityGroupId,
    });
  }
}
