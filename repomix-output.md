This file is a merged representation of the entire codebase, combined into a single document by Repomix.

# File Summary

## Purpose
This file contains a packed representation of the entire repository's contents.
It is designed to be easily consumable by AI systems for analysis, code review,
or other automated processes.

## File Format
The content is organized as follows:
1. This summary section
2. Repository information
3. Directory structure
4. Repository files (if enabled)
5. Multiple file entries, each consisting of:
  a. A header with the file path (## File: path/to/file)
  b. The full contents of the file in a code block

## Usage Guidelines
- This file should be treated as read-only. Any changes should be made to the
  original repository files, not this packed version.
- When processing this file, use the file path to distinguish
  between different files in the repository.
- Be aware that this file may contain sensitive information. Handle it with
  the same level of security as you would the original repository.

## Notes
- Some files may have been excluded based on .gitignore rules and Repomix's configuration
- Binary files are not included in this packed representation. Please refer to the Repository Structure section for a complete list of file paths, including binary files
- Files matching patterns in .gitignore are excluded
- Files matching default ignore patterns are excluded
- Files are sorted by Git change count (files with more changes are at the bottom)

# Directory Structure
```
.serena/
  memories/
    project_overview.md
    style_and_conventions.md
    suggested_commands.md
    task_completion.md
  .gitignore
  project.yml
bookingapp/
  cmd/
    api/
      main.go
  internal/
    domain/
      entity/
        plan.go
        reservation.go
        user.go
      repository/
        repository.go
    infrastructure/
      db/
        models/
          user/
            user_model.go
          plan_model.go
          reservation_model.go
        migrate.go
        mysql.go
      memory/
        plan_repo_memory.go
        reservation_repo_memory.go
      repository/
        mysql/
          user/
            user_get.go
            user_repo_mysql.go
          plan_repo_mysql.go
          reservation_repo_mysql.go
    interface/
      http/
        reservation_handler.go
        user.go
    usecase/
      user/
        user_uc.go
      reservation_uc.go
      user_uc.go
  docker-compose.yml
  Dockerfile
  go.mod
  mermaid.md
  README.md
infra/
  bin/
    infra.ts
  lib/
    infra-stack.ts
  sessionmanager-bundle/
    bin/
      session-manager-plugin
    install
    LICENSE
    NOTICE
    README.md
    RELEASENOTES.md
    seelog.xml.template
    THIRD-PARTY
    VERSION
  test/
    infra.test.ts
  .gitignore
  .npmignore
  AWSCLIV2.pkg
  cdk.json
  jest.config.js
  my-ec2-key.pem
  package.json
  README.md
  sessionmanager-bundle.zip
  tsconfig.json
fix.md
repomix-output.xml
```

# Files

## File: bookingapp/Dockerfile
````
# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Install build deps
RUN apk add --no-cache git

# Copy go module manifests first and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the bookingapp binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bookingapp ./cmd/api

# Runtime stage
FROM gcr.io/distroless/base-debian12
WORKDIR /app

COPY --from=builder /app/bookingapp ./bookingapp

EXPOSE 8080
ENTRYPOINT ["./bookingapp"]
````

## File: infra/sessionmanager-bundle/install
````
#!/usr/bin/env python3
import optparse
import os
import shutil
import sys

"""
This script installs the session-manager-plugin for macOS.
The executable is installed to /usr/local/sessionmanagerplugin (default) or to an install directory provided by the user.
It also creates a symlink session-manager-plugin in the /usr/local/bin directory
"""

PLUGIN_FILE = "session-manager-plugin"
VERSION_FILE = "VERSION"
LICENSE_FILE = "LICENSE"
SEELOG_FILE = "seelog.xml.template"

INSTALL_DIR = "/usr/local/sessionmanagerplugin"
SYMLINK_NAME = "/usr/local/bin/{}".format(PLUGIN_FILE)

def create_symlink(real_location, symlink_name):
    """
    Removes a duplicate symlink if it exists and
    creates symlink from real_location to symlink_name
    """
    if os.path.isfile(symlink_name):
        print("Symlink already exists. Removing symlink from {}".format(symlink_name))
        os.remove(symlink_name)

    print("Creating Symlink from {} to {}".format(real_location, symlink_name))
    os.symlink(real_location, symlink_name)

def main():
    parser = optparse.OptionParser()
    parser.add_option("-i", "--install-dir", help="The location to install the Session Manager Plugin."
                    " The default value is {}".format(INSTALL_DIR), default=INSTALL_DIR)
    parser.add_option("-b", "--bin-location", help="If this argument is "
                      "provided, then a symlink will be created at this "
                      "location that points to the session-manager-plugin executable. "
                      "The default symlink location is {}\n"
                      "Note: The session-manager-plugin executable must be in your $PATH "
                      "to use Session Manager Plugin with AWS CLI.".format(SYMLINK_NAME), default=SYMLINK_NAME)
    options = parser.parse_args()[0]

    try:
        current_working_directory = os.path.dirname(os.path.abspath(__file__))

        current_bin_folder = os.path.join(current_working_directory, 'bin')
        install_bin_folder = os.path.join(options.install_dir, 'bin')

        if not os.path.isdir(install_bin_folder):
             print("Creating install directories: {}".format(install_bin_folder))
             os.makedirs(install_bin_folder)

        # Copy executable. Overwrites file if it exists. The basename of the file is copied
        current_bin_location = os.path.join(current_working_directory, 'bin', PLUGIN_FILE)
        shutil.copy2(current_bin_location, install_bin_folder)
        current_bin_folder = install_bin_folder

        # Copy see_log file
        seelog_location = os.path.join(current_working_directory, SEELOG_FILE)
        shutil.copy2(seelog_location, options.install_dir)

        # Copy Version File
        version_file_location = os.path.join(current_working_directory, VERSION_FILE)
        shutil.copy2(version_file_location, options.install_dir)

        # Copy License File
        license_file_location = os.path.join(current_working_directory, LICENSE_FILE)
        shutil.copy2(license_file_location, options.install_dir)

        install_bin_location = os.path.join(options.install_dir,'bin', PLUGIN_FILE)
        create_symlink(install_bin_location, options.bin_location)
        print("Installation successful!")
    except:
       print("Failed to create symlink.\nPlease add {} to your $PATH to use Session Manager Plugin.".format(current_bin_folder))

if __name__ == '__main__':
    main()
````

## File: infra/sessionmanager-bundle/LICENSE
````
Apache License
                           Version 2.0, January 2004
                        http://www.apache.org/licenses/

   TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION

   1. Definitions.

      "License" shall mean the terms and conditions for use, reproduction,
      and distribution as defined by Sections 1 through 9 of this document.

      "Licensor" shall mean the copyright owner or entity authorized by
      the copyright owner that is granting the License.

      "Legal Entity" shall mean the union of the acting entity and all
      other entities that control, are controlled by, or are under common
      control with that entity. For the purposes of this definition,
      "control" means (i) the power, direct or indirect, to cause the
      direction or management of such entity, whether by contract or
      otherwise, or (ii) ownership of fifty percent (50%) or more of the
      outstanding shares, or (iii) beneficial ownership of such entity.

      "You" (or "Your") shall mean an individual or Legal Entity
      exercising permissions granted by this License.

      "Source" form shall mean the preferred form for making modifications,
      including but not limited to software source code, documentation
      source, and configuration files.

      "Object" form shall mean any form resulting from mechanical
      transformation or translation of a Source form, including but
      not limited to compiled object code, generated documentation,
      and conversions to other media types.

      "Work" shall mean the work of authorship, whether in Source or
      Object form, made available under the License, as indicated by a
      copyright notice that is included in or attached to the work
      (an example is provided in the Appendix below).

      "Derivative Works" shall mean any work, whether in Source or Object
      form, that is based on (or derived from) the Work and for which the
      editorial revisions, annotations, elaborations, or other modifications
      represent, as a whole, an original work of authorship. For the purposes
      of this License, Derivative Works shall not include works that remain
      separable from, or merely link (or bind by name) to the interfaces of,
      the Work and Derivative Works thereof.

      "Contribution" shall mean any work of authorship, including
      the original version of the Work and any modifications or additions
      to that Work or Derivative Works thereof, that is intentionally
      submitted to Licensor for inclusion in the Work by the copyright owner
      or by an individual or Legal Entity authorized to submit on behalf of
      the copyright owner. For the purposes of this definition, "submitted"
      means any form of electronic, verbal, or written communication sent
      to the Licensor or its representatives, including but not limited to
      communication on electronic mailing lists, source code control systems,
      and issue tracking systems that are managed by, or on behalf of, the
      Licensor for the purpose of discussing and improving the Work, but
      excluding communication that is conspicuously marked or otherwise
      designated in writing by the copyright owner as "Not a Contribution."

      "Contributor" shall mean Licensor and any individual or Legal Entity
      on behalf of whom a Contribution has been received by Licensor and
      subsequently incorporated within the Work.

   2. Grant of Copyright License. Subject to the terms and conditions of
      this License, each Contributor hereby grants to You a perpetual,
      worldwide, non-exclusive, no-charge, royalty-free, irrevocable
      copyright license to reproduce, prepare Derivative Works of,
      publicly display, publicly perform, sublicense, and distribute the
      Work and such Derivative Works in Source or Object form.

   3. Grant of Patent License. Subject to the terms and conditions of
      this License, each Contributor hereby grants to You a perpetual,
      worldwide, non-exclusive, no-charge, royalty-free, irrevocable
      (except as stated in this section) patent license to make, have made,
      use, offer to sell, sell, import, and otherwise transfer the Work,
      where such license applies only to those patent claims licensable
      by such Contributor that are necessarily infringed by their
      Contribution(s) alone or by combination of their Contribution(s)
      with the Work to which such Contribution(s) was submitted. If You
      institute patent litigation against any entity (including a
      cross-claim or counterclaim in a lawsuit) alleging that the Work
      or a Contribution incorporated within the Work constitutes direct
      or contributory patent infringement, then any patent licenses
      granted to You under this License for that Work shall terminate
      as of the date such litigation is filed.

   4. Redistribution. You may reproduce and distribute copies of the
      Work or Derivative Works thereof in any medium, with or without
      modifications, and in Source or Object form, provided that You
      meet the following conditions:

      (a) You must give any other recipients of the Work or
          Derivative Works a copy of this License; and

      (b) You must cause any modified files to carry prominent notices
          stating that You changed the files; and

      (c) You must retain, in the Source form of any Derivative Works
          that You distribute, all copyright, patent, trademark, and
          attribution notices from the Source form of the Work,
          excluding those notices that do not pertain to any part of
          the Derivative Works; and

      (d) If the Work includes a "NOTICE" text file as part of its
          distribution, then any Derivative Works that You distribute must
          include a readable copy of the attribution notices contained
          within such NOTICE file, excluding those notices that do not
          pertain to any part of the Derivative Works, in at least one
          of the following places: within a NOTICE text file distributed
          as part of the Derivative Works; within the Source form or
          documentation, if provided along with the Derivative Works; or,
          within a display generated by the Derivative Works, if and
          wherever such third-party notices normally appear. The contents
          of the NOTICE file are for informational purposes only and
          do not modify the License. You may add Your own attribution
          notices within Derivative Works that You distribute, alongside
          or as an addendum to the NOTICE text from the Work, provided
          that such additional attribution notices cannot be construed
          as modifying the License.

      You may add Your own copyright statement to Your modifications and
      may provide additional or different license terms and conditions
      for use, reproduction, or distribution of Your modifications, or
      for any such Derivative Works as a whole, provided Your use,
      reproduction, and distribution of the Work otherwise complies with
      the conditions stated in this License.

   5. Submission of Contributions. Unless You explicitly state otherwise,
      any Contribution intentionally submitted for inclusion in the Work
      by You to the Licensor shall be under the terms and conditions of
      this License, without any additional terms or conditions.
      Notwithstanding the above, nothing herein shall supersede or modify
      the terms of any separate license agreement you may have executed
      with Licensor regarding such Contributions.

   6. Trademarks. This License does not grant permission to use the trade
      names, trademarks, service marks, or product names of the Licensor,
      except as required for reasonable and customary use in describing the
      origin of the Work and reproducing the content of the NOTICE file.

   7. Disclaimer of Warranty. Unless required by applicable law or
      agreed to in writing, Licensor provides the Work (and each
      Contributor provides its Contributions) on an "AS IS" BASIS,
      WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
      implied, including, without limitation, any warranties or conditions
      of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A
      PARTICULAR PURPOSE. You are solely responsible for determining the
      appropriateness of using or redistributing the Work and assume any
      risks associated with Your exercise of permissions under this License.

   8. Limitation of Liability. In no event and under no legal theory,
      whether in tort (including negligence), contract, or otherwise,
      unless required by applicable law (such as deliberate and grossly
      negligent acts) or agreed to in writing, shall any Contributor be
      liable to You for damages, including any direct, indirect, special,
      incidental, or consequential damages of any character arising as a
      result of this License or out of the use or inability to use the
      Work (including but not limited to damages for loss of goodwill,
      work stoppage, computer failure or malfunction, or any and all
      other commercial damages or losses), even if such Contributor
      has been advised of the possibility of such damages.

   9. Accepting Warranty or Additional Liability. While redistributing
      the Work or Derivative Works thereof, You may choose to offer,
      and charge a fee for, acceptance of support, warranty, indemnity,
      or other liability obligations and/or rights consistent with this
      License. However, in accepting such obligations, You may act only
      on Your own behalf and on Your sole responsibility, not on behalf
      of any other Contributor, and only if You agree to indemnify,
      defend, and hold each Contributor harmless for any liability
      incurred by, or claims asserted against, such Contributor by reason
      of your accepting any such warranty or additional liability.

   END OF TERMS AND CONDITIONS

   APPENDIX: How to apply the Apache License to your work.

      To apply the Apache License to your work, attach the following
      boilerplate notice, with the fields enclosed by brackets "[]"
      replaced with your own identifying information. (Don't include
      the brackets!)  The text should be enclosed in the appropriate
      comment syntax for the file format. We also recommend that a
      file or class name and description of purpose be included on the
      same "printed page" as the copyright notice for easier
      identification within third-party archives.

   Copyright [yyyy] [name of copyright owner]

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
````

## File: infra/sessionmanager-bundle/NOTICE
````
Session Manager Plugin
Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
````

## File: infra/sessionmanager-bundle/README.md
````markdown
# Session Manager Plugin

This plugin helps you to use the AWS Command Line Interface (AWS CLI) to start and end sessions to your managed instances. Session Manager is a capability of AWS Systems Manager.

## Overview

Session Manager is a fully managed AWS Systems Manager capability that lets you manage your Amazon Elastic Compute Cloud (Amazon EC2) instances, on-premises instances and virtual machines. Session Manager provides secure and auditable instance management without the need to open inbound ports. When you use the Session Manager plugin with the AWS CLI to start a session, the plugin builds the websocket connection to your managed instances.

### Prerequisites

Before using Session Manager, make sure your environment meets the following requirements. [Complete Session Manager prerequisites](http://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-prerequisites.html).

### Starting a session

For information about starting a session using the AWS CLI, see [Starting a session (AWS CLI)](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-sessions-start.html#sessions-start-cli).

### Troubleshooting

For information about troubleshooting, see [Troubleshooting Session Manager](http://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-troubleshooting.html).

### Version Compatibility

The default compiled version is 1.2.0.0. We strongly recommend you manually update to version 1.3.0.0 before you compile. This step is crucial to ensure you have access to the latest features and functionality.

To update the version:
1. Locate the version information in the following files:
    - `src/version/version.go`
    - `VERSION`
2. In both files, update the version number from 1.2.0.0 to 1.3.0.0.
3. Save the changes and recompile the plugin with the updated version number.

By taking this extra step, you'll avoid potential feature limitations or functionality issues that may arise from using an outdated version. This practice ensures that you're working with the most up-to-date capabilities of the Session Manager Plugin.

### Working with Docker

To build the Session Manager plugin in a `Docker` container, complete the following steps:

1. Install [`docker`](https://docs.docker.com/engine/install/centos/)

2. Build the `docker` image
```
docker build -t session-manager-plugin-image .
```
3. Build the plugin
```
docker run -it --rm --name session-manager-plugin -v `pwd`:/session-manager-plugin session-manager-plugin-image make release
```

### Working with Linux

To build the binaries required to install the Session Manager plugin, complete the following steps.

1. Install `golang`

2. Install `rpm-build` and `rpmdevtools`

3. Install `gcc 8.3+` and `glibc 2.27+`

4. Run `make release` to build the plugin for Linux, Debian, macOS and Windows.

5. Change to the directory of your local machine's operating system architecture and open the `session-manager-plugin` directory. Then follow the installation procedure that applies to your local machine. For more information, see [Install the Session Manager plugin for the AWS CLI](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html). If the machine you're building the plugin on differs from the machine you plan to install the plugin on you will need to copy the `session-manager-plugin` binary to the appropriate directory for that operating system.

```
Linux - /usr/local/sessionmanagerplugin/bin/session-manager-plugin

macOS - /usr/local/sessionmanagerplugin/bin/session-manager-plugin

Windows - C:\Program Files\Amazon\SessionManagerPlugin\bin\session-manager-plugin.exe
```

The `ssmcli` binary is available for some operating systems for testing purposes only. The following is an example command using this binary.

```
./ssmcli start-session --instance-id i-1234567890abcdef0 --region us-east-2
```

### Directory structure

Source code

* `sessionmanagerplugin/session` contains the source code for core functionalities
* `communicator/` contains the source code for websocket related operations
* `vendor/src` contains the vendor package source code
* `packaging/` contains rpm and dpkg artifacts
* `Tools/src` contains build scripts

## Feedback

Thank you for helping us to improve the Session Manager plugin. Please send your questions or comments to the [Systems Manager Forum](https://forums.aws.amazon.com/forum.jspa?forumID=185&start=0)

## License

The session-manager-plugin is licensed under the Apache 2.0 License.
````

## File: infra/sessionmanager-bundle/RELEASENOTES.md
````markdown
Latest
================
- Update Readme for version configuration
- Upgrade Go version to 1.23 in Dockerfile

1.2.694.0
================
- Rollback change on adding credential to OpenDataChannel request

1.2.688.0
================
- Add credential to OpenDataChannel request
- Upgrade dependent packages testify-1.9.0 and objx-0.5.2

1.2.677.0
================
- Support passing plugin version with OpenDataChannel request

1.2.650.0
================
- Upgrade aws-sdk-go to 1.54.10
- Reformat comments for gofmt check

1.2.633.0
================
- Update dockerfile to use ecr image

1.2.553.0
================
- Upgrade aws-sdk-go and dependent Golang packages

1.2.536.0
================
- Support reading StartSession API response parameters from environment variable
- Migrate to use images from ECR repository

1.2.497.0
================
- Upgrade Go SDK to v1.44.302

1.2.463.0
================
- Support ARM64 for Apple Mac M1
- Remove unused start/stop steps in packaging scripts
	  
1.2.398.0
================
- Support golang version 1.17
- Update default session-manager-plugin runner for macOS to be python3
- Update import path from SSMCLI to session-manager-plugin
 
1.2.339.0
================
- Fix idle session timeout for port sessions

1.2.331.0
================
- Fix port session premature close when local server is not connected before timeout

1.2.323.0
================
- Disable smux keep alive to use idle session timeout feature

1.2.312.0
================
- Support more output message payload type.

1.2.295.0
================
- Fix hung sessions caused by client resending stream data when agent becomes inactive.
- Fix incorrect logs for start_publication and pause_publication messages.

1.2.279.0
================
- Fix single memory reference access for parameters
- Upgrade gorilla package to 1.4.2 

1.2.245.0
================
- Enhancement: Upgrade aws-sdk-go to latest version (v1.40.17) to support SSO
- Enhancement: Improve error message for legacy CLI version 

1.2.234.0
================
- Change data streaming related logs from debug to trace level.
- Fix typo for log config file.
- Fix interactive command session abruptly terminated issue.

1.2.205.0
================
- Introduce client timeout for session start request.
- Add support for signed session-manager-plugin.pkg file for macOS.

1.2.54.0
================
- Enhancement: Added support  for running session in NonInteractiveCommands execution mode.

1.2.30.0
================
- Bug Fix: (Port forwarding sessions only) Using system tmp folder for unix socket path.

1.2.7.0
================
- Enhancement: (Port forwarding sessions only) Reduced latency and improved overall performance.

1.1.61.0
================
- Enhancement: Added ARM support for Linux and Ubuntu.

1.1.54.0
================
- Bug Fix: Handle race condition scenario of packets being dropped when plugin is not ready.

1.1.50.0
================
- Enhancement: Add support for forwarding port session to local unix socket.

1.1.35.0
================
- Enhancement: For port forwarding session, send terminateSession flag to SSM agent on receiving Control-C signal.

1.1.33.0
================
- Enhancement: For port forwarding session, send disconnect flag to server when client drops tcp connection.

1.1.31.0
================
- Enhancement: Change to keep port forwarding session open until remote server closes the connection.

1.1.26.0
================
- Enhancement: Limit the rate of data transfer in port session.

1.1.23.0
================
- Enhancement: Add support for running SSH sessions using Session Manager.

1.1.17.0
================
- Enhancement: Add support for further encryption of session data using AWS KMS.

1.0.37.0
================
- Fix bug for Windows SessionManagerPlugin

1.0.0.0
================
- Initial SessionManagerPlugin release
````

## File: infra/sessionmanager-bundle/seelog.xml.template
````
<!--SessionManagerPlugin uses seelog logging -->
<!--Seelog has github wiki pages, which contain detailed how-tos references: https://github.com/cihub/seelog/wiki -->
<!--Seelog examples can be found here: https://github.com/cihub/seelog-examples -->
<seelog type="adaptive" mininterval="2000000" maxinterval="100000000" critmsgcount="500" minlevel="off">
    <exceptions>
        <exception filepattern="test*" minlevel="error"/>
    </exceptions>
    <outputs formatid="fmtinfo">
        <rollingfile type="size" filename="/usr/local/sessionmanagerplugin/logs/session-manager-plugin.log" maxsize="30000000" maxrolls="5"/>
        <filter levels="error,critical" formatid="fmterror">
            <rollingfile type="size" filename="/usr/local/sessionmanagerplugin/logs/errors.log" maxsize="10000000" maxrolls="5"/>
        </filter>
    </outputs>
    <formats>
        <format id="fmterror" format="%Date %Time %LEVEL [%FuncShort @ %File.%Line] %Msg%n"/>
        <format id="fmtdebug" format="%Date %Time %LEVEL [%FuncShort @ %File.%Line] %Msg%n"/>
        <format id="fmtinfo" format="%Date %Time %LEVEL %Msg%n"/>
    </formats>
</seelog>
````

## File: infra/sessionmanager-bundle/THIRD-PARTY
````
The Amazon Session Manager Plugin constitutes AWS Content as defined in the AWS Customer Agreement
or your relevant customer agreement with AWS, and is licensed to you under that agreement.

The Amazon Session Manager Plugin includes the following third-party software/licensing:

** cihub/seelog - https://github.com/cihub/seelog
Copyright (c) 2012, Cloud Instruments Co., Ltd. <info@cin.io>. All rights reserved.
** gorilla/websocket - https://github.com/gorilla/websocket
Copyright (c) 2013 The Gorilla WebSocket Authors. All rights reserved.
** fsnotigy/fsnotify - https://github.com/fsnotify/fsnotify
Copyright (c) 2012 The Go Authors. All rights reserved.
Copyright (c) 2012 fsnotify Authors. All rights reserved.
** pmezard/go-difflib - https://github.com/pmezard/go-difflib
Copyright (c) 2013, Patrick Mezard

BSD License

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:
    * Redistributions of source code must retain the above copyright
      notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above copyright
      notice, this list of conditions and the following disclaimer in the
      documentation and/or other materials provided with the distribution.
    * Neither the name of the Cloud Instruments Co., Ltd. nor the
      names of its contributors may be used to endorse or promote products
      derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

----------------

** twinj/uuid - https://github.com/twinj/uuid
Copyright (C) 2011 by Krzysztof Kowalik <chris@nu7hat.ch>
Copyright (C) 2016 by Daniel Kemp <twinj@github.com> Derivative work
** stretchr/testify - https://github.com/stretchr/testify
Copyright (c) 2012 - 2013 Mat Ryer and Tyler Bunnell
** stretchr/objx - https://github.com/stretchr/objx
Copyright (c) 2014 Stretchr, Inc.
Copyright (c) 2017-2018 objx contributors
** eiannone/keyboard - https://github.com/eiannone/keyboard
Copyright (C) 2012 termbox-go authors
Copyright (c) 2015 Emanuele Iannone
** xtaci/smux - https://github.com/xtaci/smux
Copyright (c) 2016-2017 Daniel Fu

MIT License

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

----------------

** davecgh/go-spew - https://github.com/davecgh/go-spew
Copyright (c) 2012-2016 Dave Collins <dave@davec.name>

ISC License

Permission to use, copy, modify, and/or distribute this software for any
purpose with or without fee is hereby granted, provided that the above
copyright notice and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

----------------

** jmespath/go-jmespath - https://github.com/jmespath/go-jmespath
Copyright 2015 James Saryerwinnie
** go-yaml/yaml - https://github.com/go-yaml/yaml
Copyright (c) 2011-2019 Canonical Ltd

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
````

## File: infra/sessionmanager-bundle/VERSION
````
1.2.707.0
````

## File: repomix-output.xml
````xml
This file is a merged representation of the entire codebase, combined into a single document by Repomix.

<file_summary>
This section contains a summary of this file.

<purpose>
This file contains a packed representation of the entire repository's contents.
It is designed to be easily consumable by AI systems for analysis, code review,
or other automated processes.
</purpose>

<file_format>
The content is organized as follows:
1. This summary section
2. Repository information
3. Directory structure
4. Repository files (if enabled)
5. Multiple file entries, each consisting of:
  - File path as an attribute
  - Full contents of the file
</file_format>

<usage_guidelines>
- This file should be treated as read-only. Any changes should be made to the
  original repository files, not this packed version.
- When processing this file, use the file path to distinguish
  between different files in the repository.
- Be aware that this file may contain sensitive information. Handle it with
  the same level of security as you would the original repository.
</usage_guidelines>

<notes>
- Some files may have been excluded based on .gitignore rules and Repomix's configuration
- Binary files are not included in this packed representation. Please refer to the Repository Structure section for a complete list of file paths, including binary files
- Files matching patterns in .gitignore are excluded
- Files matching default ignore patterns are excluded
- Files are sorted by Git change count (files with more changes are at the bottom)
</notes>

</file_summary>

<directory_structure>
.serena/
  memories/
    project_overview.md
    style_and_conventions.md
    suggested_commands.md
    task_completion.md
  .gitignore
  project.yml
bookingapp/
  cmd/
    api/
      main.go
  internal/
    domain/
      entity/
        plan.go
        reservation.go
        user.go
      repository/
        repository.go
    infrastructure/
      db/
        models/
          user/
            user_model.go
          plan_model.go
          reservation_model.go
        migrate.go
        mysql.go
      memory/
        plan_repo_memory.go
        reservation_repo_memory.go
      repository/
        mysql/
          user/
            user_get.go
            user_repo_mysql.go
          plan_repo_mysql.go
          reservation_repo_mysql.go
    interface/
      http/
        reservation_handler.go
        user.go
    usecase/
      user/
        user_uc.go
      reservation_uc.go
      user_uc.go
  docker-compose.yml
  Dockerfile
  go.mod
  mermaid.md
  README.md
infra/
  bin/
    infra.ts
  lib/
    infra-stack.ts
  sessionmanager-bundle/
    bin/
      session-manager-plugin
    install
    LICENSE
    NOTICE
    README.md
    RELEASENOTES.md
    seelog.xml.template
    THIRD-PARTY
    VERSION
  test/
    infra.test.ts
  .gitignore
  .npmignore
  AWSCLIV2.pkg
  cdk.json
  jest.config.js
  my-ec2-key.pem
  package.json
  README.md
  sessionmanager-bundle.zip
  tsconfig.json
fix.md
</directory_structure>

<files>
This section contains the contents of the repository's files.

<file path="bookingapp/Dockerfile">
# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Install build deps
RUN apk add --no-cache git

# Copy go module manifests first and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the bookingapp binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bookingapp ./cmd/api

# Runtime stage
FROM gcr.io/distroless/base-debian12
WORKDIR /app

COPY --from=builder /app/bookingapp ./bookingapp

EXPOSE 8080
ENTRYPOINT ["./bookingapp"]
</file>

<file path="infra/sessionmanager-bundle/install">
#!/usr/bin/env python3
import optparse
import os
import shutil
import sys

"""
This script installs the session-manager-plugin for macOS.
The executable is installed to /usr/local/sessionmanagerplugin (default) or to an install directory provided by the user.
It also creates a symlink session-manager-plugin in the /usr/local/bin directory
"""

PLUGIN_FILE = "session-manager-plugin"
VERSION_FILE = "VERSION"
LICENSE_FILE = "LICENSE"
SEELOG_FILE = "seelog.xml.template"

INSTALL_DIR = "/usr/local/sessionmanagerplugin"
SYMLINK_NAME = "/usr/local/bin/{}".format(PLUGIN_FILE)

def create_symlink(real_location, symlink_name):
    """
    Removes a duplicate symlink if it exists and
    creates symlink from real_location to symlink_name
    """
    if os.path.isfile(symlink_name):
        print("Symlink already exists. Removing symlink from {}".format(symlink_name))
        os.remove(symlink_name)

    print("Creating Symlink from {} to {}".format(real_location, symlink_name))
    os.symlink(real_location, symlink_name)

def main():
    parser = optparse.OptionParser()
    parser.add_option("-i", "--install-dir", help="The location to install the Session Manager Plugin."
                    " The default value is {}".format(INSTALL_DIR), default=INSTALL_DIR)
    parser.add_option("-b", "--bin-location", help="If this argument is "
                      "provided, then a symlink will be created at this "
                      "location that points to the session-manager-plugin executable. "
                      "The default symlink location is {}\n"
                      "Note: The session-manager-plugin executable must be in your $PATH "
                      "to use Session Manager Plugin with AWS CLI.".format(SYMLINK_NAME), default=SYMLINK_NAME)
    options = parser.parse_args()[0]

    try:
        current_working_directory = os.path.dirname(os.path.abspath(__file__))

        current_bin_folder = os.path.join(current_working_directory, 'bin')
        install_bin_folder = os.path.join(options.install_dir, 'bin')

        if not os.path.isdir(install_bin_folder):
             print("Creating install directories: {}".format(install_bin_folder))
             os.makedirs(install_bin_folder)

        # Copy executable. Overwrites file if it exists. The basename of the file is copied
        current_bin_location = os.path.join(current_working_directory, 'bin', PLUGIN_FILE)
        shutil.copy2(current_bin_location, install_bin_folder)
        current_bin_folder = install_bin_folder

        # Copy see_log file
        seelog_location = os.path.join(current_working_directory, SEELOG_FILE)
        shutil.copy2(seelog_location, options.install_dir)

        # Copy Version File
        version_file_location = os.path.join(current_working_directory, VERSION_FILE)
        shutil.copy2(version_file_location, options.install_dir)

        # Copy License File
        license_file_location = os.path.join(current_working_directory, LICENSE_FILE)
        shutil.copy2(license_file_location, options.install_dir)

        install_bin_location = os.path.join(options.install_dir,'bin', PLUGIN_FILE)
        create_symlink(install_bin_location, options.bin_location)
        print("Installation successful!")
    except:
       print("Failed to create symlink.\nPlease add {} to your $PATH to use Session Manager Plugin.".format(current_bin_folder))

if __name__ == '__main__':
    main()
</file>

<file path="infra/sessionmanager-bundle/LICENSE">
Apache License
                           Version 2.0, January 2004
                        http://www.apache.org/licenses/

   TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION

   1. Definitions.

      "License" shall mean the terms and conditions for use, reproduction,
      and distribution as defined by Sections 1 through 9 of this document.

      "Licensor" shall mean the copyright owner or entity authorized by
      the copyright owner that is granting the License.

      "Legal Entity" shall mean the union of the acting entity and all
      other entities that control, are controlled by, or are under common
      control with that entity. For the purposes of this definition,
      "control" means (i) the power, direct or indirect, to cause the
      direction or management of such entity, whether by contract or
      otherwise, or (ii) ownership of fifty percent (50%) or more of the
      outstanding shares, or (iii) beneficial ownership of such entity.

      "You" (or "Your") shall mean an individual or Legal Entity
      exercising permissions granted by this License.

      "Source" form shall mean the preferred form for making modifications,
      including but not limited to software source code, documentation
      source, and configuration files.

      "Object" form shall mean any form resulting from mechanical
      transformation or translation of a Source form, including but
      not limited to compiled object code, generated documentation,
      and conversions to other media types.

      "Work" shall mean the work of authorship, whether in Source or
      Object form, made available under the License, as indicated by a
      copyright notice that is included in or attached to the work
      (an example is provided in the Appendix below).

      "Derivative Works" shall mean any work, whether in Source or Object
      form, that is based on (or derived from) the Work and for which the
      editorial revisions, annotations, elaborations, or other modifications
      represent, as a whole, an original work of authorship. For the purposes
      of this License, Derivative Works shall not include works that remain
      separable from, or merely link (or bind by name) to the interfaces of,
      the Work and Derivative Works thereof.

      "Contribution" shall mean any work of authorship, including
      the original version of the Work and any modifications or additions
      to that Work or Derivative Works thereof, that is intentionally
      submitted to Licensor for inclusion in the Work by the copyright owner
      or by an individual or Legal Entity authorized to submit on behalf of
      the copyright owner. For the purposes of this definition, "submitted"
      means any form of electronic, verbal, or written communication sent
      to the Licensor or its representatives, including but not limited to
      communication on electronic mailing lists, source code control systems,
      and issue tracking systems that are managed by, or on behalf of, the
      Licensor for the purpose of discussing and improving the Work, but
      excluding communication that is conspicuously marked or otherwise
      designated in writing by the copyright owner as "Not a Contribution."

      "Contributor" shall mean Licensor and any individual or Legal Entity
      on behalf of whom a Contribution has been received by Licensor and
      subsequently incorporated within the Work.

   2. Grant of Copyright License. Subject to the terms and conditions of
      this License, each Contributor hereby grants to You a perpetual,
      worldwide, non-exclusive, no-charge, royalty-free, irrevocable
      copyright license to reproduce, prepare Derivative Works of,
      publicly display, publicly perform, sublicense, and distribute the
      Work and such Derivative Works in Source or Object form.

   3. Grant of Patent License. Subject to the terms and conditions of
      this License, each Contributor hereby grants to You a perpetual,
      worldwide, non-exclusive, no-charge, royalty-free, irrevocable
      (except as stated in this section) patent license to make, have made,
      use, offer to sell, sell, import, and otherwise transfer the Work,
      where such license applies only to those patent claims licensable
      by such Contributor that are necessarily infringed by their
      Contribution(s) alone or by combination of their Contribution(s)
      with the Work to which such Contribution(s) was submitted. If You
      institute patent litigation against any entity (including a
      cross-claim or counterclaim in a lawsuit) alleging that the Work
      or a Contribution incorporated within the Work constitutes direct
      or contributory patent infringement, then any patent licenses
      granted to You under this License for that Work shall terminate
      as of the date such litigation is filed.

   4. Redistribution. You may reproduce and distribute copies of the
      Work or Derivative Works thereof in any medium, with or without
      modifications, and in Source or Object form, provided that You
      meet the following conditions:

      (a) You must give any other recipients of the Work or
          Derivative Works a copy of this License; and

      (b) You must cause any modified files to carry prominent notices
          stating that You changed the files; and

      (c) You must retain, in the Source form of any Derivative Works
          that You distribute, all copyright, patent, trademark, and
          attribution notices from the Source form of the Work,
          excluding those notices that do not pertain to any part of
          the Derivative Works; and

      (d) If the Work includes a "NOTICE" text file as part of its
          distribution, then any Derivative Works that You distribute must
          include a readable copy of the attribution notices contained
          within such NOTICE file, excluding those notices that do not
          pertain to any part of the Derivative Works, in at least one
          of the following places: within a NOTICE text file distributed
          as part of the Derivative Works; within the Source form or
          documentation, if provided along with the Derivative Works; or,
          within a display generated by the Derivative Works, if and
          wherever such third-party notices normally appear. The contents
          of the NOTICE file are for informational purposes only and
          do not modify the License. You may add Your own attribution
          notices within Derivative Works that You distribute, alongside
          or as an addendum to the NOTICE text from the Work, provided
          that such additional attribution notices cannot be construed
          as modifying the License.

      You may add Your own copyright statement to Your modifications and
      may provide additional or different license terms and conditions
      for use, reproduction, or distribution of Your modifications, or
      for any such Derivative Works as a whole, provided Your use,
      reproduction, and distribution of the Work otherwise complies with
      the conditions stated in this License.

   5. Submission of Contributions. Unless You explicitly state otherwise,
      any Contribution intentionally submitted for inclusion in the Work
      by You to the Licensor shall be under the terms and conditions of
      this License, without any additional terms or conditions.
      Notwithstanding the above, nothing herein shall supersede or modify
      the terms of any separate license agreement you may have executed
      with Licensor regarding such Contributions.

   6. Trademarks. This License does not grant permission to use the trade
      names, trademarks, service marks, or product names of the Licensor,
      except as required for reasonable and customary use in describing the
      origin of the Work and reproducing the content of the NOTICE file.

   7. Disclaimer of Warranty. Unless required by applicable law or
      agreed to in writing, Licensor provides the Work (and each
      Contributor provides its Contributions) on an "AS IS" BASIS,
      WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
      implied, including, without limitation, any warranties or conditions
      of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A
      PARTICULAR PURPOSE. You are solely responsible for determining the
      appropriateness of using or redistributing the Work and assume any
      risks associated with Your exercise of permissions under this License.

   8. Limitation of Liability. In no event and under no legal theory,
      whether in tort (including negligence), contract, or otherwise,
      unless required by applicable law (such as deliberate and grossly
      negligent acts) or agreed to in writing, shall any Contributor be
      liable to You for damages, including any direct, indirect, special,
      incidental, or consequential damages of any character arising as a
      result of this License or out of the use or inability to use the
      Work (including but not limited to damages for loss of goodwill,
      work stoppage, computer failure or malfunction, or any and all
      other commercial damages or losses), even if such Contributor
      has been advised of the possibility of such damages.

   9. Accepting Warranty or Additional Liability. While redistributing
      the Work or Derivative Works thereof, You may choose to offer,
      and charge a fee for, acceptance of support, warranty, indemnity,
      or other liability obligations and/or rights consistent with this
      License. However, in accepting such obligations, You may act only
      on Your own behalf and on Your sole responsibility, not on behalf
      of any other Contributor, and only if You agree to indemnify,
      defend, and hold each Contributor harmless for any liability
      incurred by, or claims asserted against, such Contributor by reason
      of your accepting any such warranty or additional liability.

   END OF TERMS AND CONDITIONS

   APPENDIX: How to apply the Apache License to your work.

      To apply the Apache License to your work, attach the following
      boilerplate notice, with the fields enclosed by brackets "[]"
      replaced with your own identifying information. (Don't include
      the brackets!)  The text should be enclosed in the appropriate
      comment syntax for the file format. We also recommend that a
      file or class name and description of purpose be included on the
      same "printed page" as the copyright notice for easier
      identification within third-party archives.

   Copyright [yyyy] [name of copyright owner]

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
</file>

<file path="infra/sessionmanager-bundle/NOTICE">
Session Manager Plugin
Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
</file>

<file path="infra/sessionmanager-bundle/README.md">
# Session Manager Plugin

This plugin helps you to use the AWS Command Line Interface (AWS CLI) to start and end sessions to your managed instances. Session Manager is a capability of AWS Systems Manager.

## Overview

Session Manager is a fully managed AWS Systems Manager capability that lets you manage your Amazon Elastic Compute Cloud (Amazon EC2) instances, on-premises instances and virtual machines. Session Manager provides secure and auditable instance management without the need to open inbound ports. When you use the Session Manager plugin with the AWS CLI to start a session, the plugin builds the websocket connection to your managed instances.

### Prerequisites

Before using Session Manager, make sure your environment meets the following requirements. [Complete Session Manager prerequisites](http://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-prerequisites.html).

### Starting a session

For information about starting a session using the AWS CLI, see [Starting a session (AWS CLI)](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-sessions-start.html#sessions-start-cli).

### Troubleshooting

For information about troubleshooting, see [Troubleshooting Session Manager](http://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-troubleshooting.html).

### Version Compatibility

The default compiled version is 1.2.0.0. We strongly recommend you manually update to version 1.3.0.0 before you compile. This step is crucial to ensure you have access to the latest features and functionality.

To update the version:
1. Locate the version information in the following files:
    - `src/version/version.go`
    - `VERSION`
2. In both files, update the version number from 1.2.0.0 to 1.3.0.0.
3. Save the changes and recompile the plugin with the updated version number.

By taking this extra step, you'll avoid potential feature limitations or functionality issues that may arise from using an outdated version. This practice ensures that you're working with the most up-to-date capabilities of the Session Manager Plugin.

### Working with Docker

To build the Session Manager plugin in a `Docker` container, complete the following steps:

1. Install [`docker`](https://docs.docker.com/engine/install/centos/)

2. Build the `docker` image
```
docker build -t session-manager-plugin-image .
```
3. Build the plugin
```
docker run -it --rm --name session-manager-plugin -v `pwd`:/session-manager-plugin session-manager-plugin-image make release
```

### Working with Linux

To build the binaries required to install the Session Manager plugin, complete the following steps.

1. Install `golang`

2. Install `rpm-build` and `rpmdevtools`

3. Install `gcc 8.3+` and `glibc 2.27+`

4. Run `make release` to build the plugin for Linux, Debian, macOS and Windows.

5. Change to the directory of your local machine's operating system architecture and open the `session-manager-plugin` directory. Then follow the installation procedure that applies to your local machine. For more information, see [Install the Session Manager plugin for the AWS CLI](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html). If the machine you're building the plugin on differs from the machine you plan to install the plugin on you will need to copy the `session-manager-plugin` binary to the appropriate directory for that operating system.

```
Linux - /usr/local/sessionmanagerplugin/bin/session-manager-plugin

macOS - /usr/local/sessionmanagerplugin/bin/session-manager-plugin

Windows - C:\Program Files\Amazon\SessionManagerPlugin\bin\session-manager-plugin.exe
```

The `ssmcli` binary is available for some operating systems for testing purposes only. The following is an example command using this binary.

```
./ssmcli start-session --instance-id i-1234567890abcdef0 --region us-east-2
```

### Directory structure

Source code

* `sessionmanagerplugin/session` contains the source code for core functionalities
* `communicator/` contains the source code for websocket related operations
* `vendor/src` contains the vendor package source code
* `packaging/` contains rpm and dpkg artifacts
* `Tools/src` contains build scripts

## Feedback

Thank you for helping us to improve the Session Manager plugin. Please send your questions or comments to the [Systems Manager Forum](https://forums.aws.amazon.com/forum.jspa?forumID=185&start=0)

## License

The session-manager-plugin is licensed under the Apache 2.0 License.
</file>

<file path="infra/sessionmanager-bundle/RELEASENOTES.md">
Latest
================
- Update Readme for version configuration
- Upgrade Go version to 1.23 in Dockerfile

1.2.694.0
================
- Rollback change on adding credential to OpenDataChannel request

1.2.688.0
================
- Add credential to OpenDataChannel request
- Upgrade dependent packages testify-1.9.0 and objx-0.5.2

1.2.677.0
================
- Support passing plugin version with OpenDataChannel request

1.2.650.0
================
- Upgrade aws-sdk-go to 1.54.10
- Reformat comments for gofmt check

1.2.633.0
================
- Update dockerfile to use ecr image

1.2.553.0
================
- Upgrade aws-sdk-go and dependent Golang packages

1.2.536.0
================
- Support reading StartSession API response parameters from environment variable
- Migrate to use images from ECR repository

1.2.497.0
================
- Upgrade Go SDK to v1.44.302

1.2.463.0
================
- Support ARM64 for Apple Mac M1
- Remove unused start/stop steps in packaging scripts
	  
1.2.398.0
================
- Support golang version 1.17
- Update default session-manager-plugin runner for macOS to be python3
- Update import path from SSMCLI to session-manager-plugin
 
1.2.339.0
================
- Fix idle session timeout for port sessions

1.2.331.0
================
- Fix port session premature close when local server is not connected before timeout

1.2.323.0
================
- Disable smux keep alive to use idle session timeout feature

1.2.312.0
================
- Support more output message payload type.

1.2.295.0
================
- Fix hung sessions caused by client resending stream data when agent becomes inactive.
- Fix incorrect logs for start_publication and pause_publication messages.

1.2.279.0
================
- Fix single memory reference access for parameters
- Upgrade gorilla package to 1.4.2 

1.2.245.0
================
- Enhancement: Upgrade aws-sdk-go to latest version (v1.40.17) to support SSO
- Enhancement: Improve error message for legacy CLI version 

1.2.234.0
================
- Change data streaming related logs from debug to trace level.
- Fix typo for log config file.
- Fix interactive command session abruptly terminated issue.

1.2.205.0
================
- Introduce client timeout for session start request.
- Add support for signed session-manager-plugin.pkg file for macOS.

1.2.54.0
================
- Enhancement: Added support  for running session in NonInteractiveCommands execution mode.

1.2.30.0
================
- Bug Fix: (Port forwarding sessions only) Using system tmp folder for unix socket path.

1.2.7.0
================
- Enhancement: (Port forwarding sessions only) Reduced latency and improved overall performance.

1.1.61.0
================
- Enhancement: Added ARM support for Linux and Ubuntu.

1.1.54.0
================
- Bug Fix: Handle race condition scenario of packets being dropped when plugin is not ready.

1.1.50.0
================
- Enhancement: Add support for forwarding port session to local unix socket.

1.1.35.0
================
- Enhancement: For port forwarding session, send terminateSession flag to SSM agent on receiving Control-C signal.

1.1.33.0
================
- Enhancement: For port forwarding session, send disconnect flag to server when client drops tcp connection.

1.1.31.0
================
- Enhancement: Change to keep port forwarding session open until remote server closes the connection.

1.1.26.0
================
- Enhancement: Limit the rate of data transfer in port session.

1.1.23.0
================
- Enhancement: Add support for running SSH sessions using Session Manager.

1.1.17.0
================
- Enhancement: Add support for further encryption of session data using AWS KMS.

1.0.37.0
================
- Fix bug for Windows SessionManagerPlugin

1.0.0.0
================
- Initial SessionManagerPlugin release
</file>

<file path="infra/sessionmanager-bundle/seelog.xml.template">
<!--SessionManagerPlugin uses seelog logging -->
<!--Seelog has github wiki pages, which contain detailed how-tos references: https://github.com/cihub/seelog/wiki -->
<!--Seelog examples can be found here: https://github.com/cihub/seelog-examples -->
<seelog type="adaptive" mininterval="2000000" maxinterval="100000000" critmsgcount="500" minlevel="off">
    <exceptions>
        <exception filepattern="test*" minlevel="error"/>
    </exceptions>
    <outputs formatid="fmtinfo">
        <rollingfile type="size" filename="/usr/local/sessionmanagerplugin/logs/session-manager-plugin.log" maxsize="30000000" maxrolls="5"/>
        <filter levels="error,critical" formatid="fmterror">
            <rollingfile type="size" filename="/usr/local/sessionmanagerplugin/logs/errors.log" maxsize="10000000" maxrolls="5"/>
        </filter>
    </outputs>
    <formats>
        <format id="fmterror" format="%Date %Time %LEVEL [%FuncShort @ %File.%Line] %Msg%n"/>
        <format id="fmtdebug" format="%Date %Time %LEVEL [%FuncShort @ %File.%Line] %Msg%n"/>
        <format id="fmtinfo" format="%Date %Time %LEVEL %Msg%n"/>
    </formats>
</seelog>
</file>

<file path="infra/sessionmanager-bundle/THIRD-PARTY">
The Amazon Session Manager Plugin constitutes AWS Content as defined in the AWS Customer Agreement
or your relevant customer agreement with AWS, and is licensed to you under that agreement.

The Amazon Session Manager Plugin includes the following third-party software/licensing:

** cihub/seelog - https://github.com/cihub/seelog
Copyright (c) 2012, Cloud Instruments Co., Ltd. <info@cin.io>. All rights reserved.
** gorilla/websocket - https://github.com/gorilla/websocket
Copyright (c) 2013 The Gorilla WebSocket Authors. All rights reserved.
** fsnotigy/fsnotify - https://github.com/fsnotify/fsnotify
Copyright (c) 2012 The Go Authors. All rights reserved.
Copyright (c) 2012 fsnotify Authors. All rights reserved.
** pmezard/go-difflib - https://github.com/pmezard/go-difflib
Copyright (c) 2013, Patrick Mezard

BSD License

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:
    * Redistributions of source code must retain the above copyright
      notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above copyright
      notice, this list of conditions and the following disclaimer in the
      documentation and/or other materials provided with the distribution.
    * Neither the name of the Cloud Instruments Co., Ltd. nor the
      names of its contributors may be used to endorse or promote products
      derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

----------------

** twinj/uuid - https://github.com/twinj/uuid
Copyright (C) 2011 by Krzysztof Kowalik <chris@nu7hat.ch>
Copyright (C) 2016 by Daniel Kemp <twinj@github.com> Derivative work
** stretchr/testify - https://github.com/stretchr/testify
Copyright (c) 2012 - 2013 Mat Ryer and Tyler Bunnell
** stretchr/objx - https://github.com/stretchr/objx
Copyright (c) 2014 Stretchr, Inc.
Copyright (c) 2017-2018 objx contributors
** eiannone/keyboard - https://github.com/eiannone/keyboard
Copyright (C) 2012 termbox-go authors
Copyright (c) 2015 Emanuele Iannone
** xtaci/smux - https://github.com/xtaci/smux
Copyright (c) 2016-2017 Daniel Fu

MIT License

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

----------------

** davecgh/go-spew - https://github.com/davecgh/go-spew
Copyright (c) 2012-2016 Dave Collins <dave@davec.name>

ISC License

Permission to use, copy, modify, and/or distribute this software for any
purpose with or without fee is hereby granted, provided that the above
copyright notice and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

----------------

** jmespath/go-jmespath - https://github.com/jmespath/go-jmespath
Copyright 2015 James Saryerwinnie
** go-yaml/yaml - https://github.com/go-yaml/yaml
Copyright (c) 2011-2019 Canonical Ltd

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
</file>

<file path="infra/sessionmanager-bundle/VERSION">
1.2.707.0
</file>

<file path=".serena/memories/project_overview.md">
# BookingApp Overview
- Purpose: Clean Architecture sample for hotel reservation bookings with RESTful HTTP API backed by MySQL via GORM.
- Tech stack: Go 1.24+ modules, net/http, GORM (mysql driver), optional in-memory repos for tests.
- Structure: `cmd/api` bootstrap, `internal/domain` entities + repository contracts, `internal/usecase` orchestrates, `internal/interface/http` HTTP handlers, `internal/infrastructure/db` GORM models + migrations, `internal/infrastructure/memory` in-memory repos.
- Key patterns: layered Clean Architecture with repositories injected into usecases, HTTP handlers consuming usecases, optional DB seeding on startup.
- Running env: expects MySQL reachable via env vars `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASS`, `DB_NAME`; seeds sample plans if table empty.
</file>

<file path=".serena/memories/style_and_conventions.md">
# Style and Conventions
- Standard Go style: gofmt formatting, short receiver names, exported types for domain entities/usecases, unexported struct fields where possible.
- Clean Architecture layering: domain entities free of infra deps; usecases depend on domain repository interfaces; interface/http packages convert transport payloads.
- Error handling: return Go errors, map to HTTP status codes in handlers; prefer sentinel errors (`ErrInvalidDates`, etc.).
- Repositories expose interfaces for Plan and Reservation storage; concrete adapters live under infrastructure (MySQL, in-memory).
- Testing not present yet; in-memory repos can support unit tests without DB.
- Comments primarily Japanese/English mix, keep brief purposeful notes when logic not obvious.
</file>

<file path=".serena/memories/suggested_commands.md">
# Suggested Commands
- `go run ./cmd/api` – start the HTTP API (uses env vars `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASS`, `DB_NAME`; defaults to localhost MySQL, auto-migrates + seeds plans).
- `go build ./cmd/api` – compile the server binary.
- `go test ./...` – run all Go unit tests (none yet, but command stays standard).
- `gofmt -w <files>` – format Go source before committing.
- `golangci-lint run` – if linting is added; current project does not vendor config.
</file>

<file path=".serena/memories/task_completion.md">
# Task Completion Checklist
- Ensure modified Go files are gofmt formatted (`gofmt -w`).
- Run `go test ./...` to confirm unit tests (once they exist) pass.
- If server behavior affected, consider `go run ./cmd/api` smoke test against local MySQL.
- Update documentation/comments when changing exposed behavior or contracts.
- Summarize changes and note any manual verification steps for the user.
</file>

<file path=".serena/.gitignore">
/cache
</file>

<file path=".serena/project.yml">
# language of the project (csharp, python, rust, java, typescript, go, cpp, or ruby)
#  * For C, use cpp
#  * For JavaScript, use typescript
# Special requirements:
#  * csharp: Requires the presence of a .sln file in the project folder.
language: go

# whether to use the project's gitignore file to ignore files
# Added on 2025-04-07
ignore_all_files_in_gitignore: true
# list of additional paths to ignore
# same syntax as gitignore, so you can use * and **
# Was previously called `ignored_dirs`, please update your config if you are using that.
# Added (renamed) on 2025-04-07
ignored_paths: []

# whether the project is in read-only mode
# If set to true, all editing tools will be disabled and attempts to use them will result in an error
# Added on 2025-04-18
read_only: false

# list of tool names to exclude. We recommend not excluding any tools, see the readme for more details.
# Below is the complete list of tools for convenience.
# To make sure you have the latest list of tools, and to view their descriptions, 
# execute `uv run scripts/print_tool_overview.py`.
#
#  * `activate_project`: Activates a project by name.
#  * `check_onboarding_performed`: Checks whether project onboarding was already performed.
#  * `create_text_file`: Creates/overwrites a file in the project directory.
#  * `delete_lines`: Deletes a range of lines within a file.
#  * `delete_memory`: Deletes a memory from Serena's project-specific memory store.
#  * `execute_shell_command`: Executes a shell command.
#  * `find_referencing_code_snippets`: Finds code snippets in which the symbol at the given location is referenced.
#  * `find_referencing_symbols`: Finds symbols that reference the symbol at the given location (optionally filtered by type).
#  * `find_symbol`: Performs a global (or local) search for symbols with/containing a given name/substring (optionally filtered by type).
#  * `get_current_config`: Prints the current configuration of the agent, including the active and available projects, tools, contexts, and modes.
#  * `get_symbols_overview`: Gets an overview of the top-level symbols defined in a given file.
#  * `initial_instructions`: Gets the initial instructions for the current project.
#     Should only be used in settings where the system prompt cannot be set,
#     e.g. in clients you have no control over, like Claude Desktop.
#  * `insert_after_symbol`: Inserts content after the end of the definition of a given symbol.
#  * `insert_at_line`: Inserts content at a given line in a file.
#  * `insert_before_symbol`: Inserts content before the beginning of the definition of a given symbol.
#  * `list_dir`: Lists files and directories in the given directory (optionally with recursion).
#  * `list_memories`: Lists memories in Serena's project-specific memory store.
#  * `onboarding`: Performs onboarding (identifying the project structure and essential tasks, e.g. for testing or building).
#  * `prepare_for_new_conversation`: Provides instructions for preparing for a new conversation (in order to continue with the necessary context).
#  * `read_file`: Reads a file within the project directory.
#  * `read_memory`: Reads the memory with the given name from Serena's project-specific memory store.
#  * `remove_project`: Removes a project from the Serena configuration.
#  * `replace_lines`: Replaces a range of lines within a file with new content.
#  * `replace_symbol_body`: Replaces the full definition of a symbol.
#  * `restart_language_server`: Restarts the language server, may be necessary when edits not through Serena happen.
#  * `search_for_pattern`: Performs a search for a pattern in the project.
#  * `summarize_changes`: Provides instructions for summarizing the changes made to the codebase.
#  * `switch_modes`: Activates modes by providing a list of their names
#  * `think_about_collected_information`: Thinking tool for pondering the completeness of collected information.
#  * `think_about_task_adherence`: Thinking tool for determining whether the agent is still on track with the current task.
#  * `think_about_whether_you_are_done`: Thinking tool for determining whether the task is truly completed.
#  * `write_memory`: Writes a named memory (for future reference) to Serena's project-specific memory store.
excluded_tools: []

# initial prompt for the project. It will always be given to the LLM upon activating the project
# (contrary to the memories, which are loaded on demand).
initial_prompt: ""

project_name: "CleanArc"
</file>

<file path="bookingapp/cmd/api/main.go">
package main

import (
	"bookingapp/internal/infrastructure/db"
	"bookingapp/internal/infrastructure/db/models"
	mysqlrepo "bookingapp/internal/infrastructure/repository/mysql"
	userrepo "bookingapp/internal/infrastructure/repository/mysql/user"
	httpi "bookingapp/internal/interface/http"
	"bookingapp/internal/usecase"
	"log"
	"net/http"
	"os"
	"strconv"

	"gorm.io/gorm"
)

func main() {
	// ---- 環境変数から接続情報 ----
	host := getEnv("DB_HOST", "127.0.0.1")
	port := getEnvInt("DB_PORT", 3306)
	user := getEnv("DB_USER", "root")
	pass := getEnv("DB_PASS", "password")
	name := getEnv("DB_NAME", "booking")

	gdb, err := db.Open(db.Config{
		User: user, Pass: pass, Host: host, Port: port, Name: name,
	})
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	if err := db.Ping(gdb); err != nil {
		log.Fatalf("ping db: %v", err)
	}
	if err := db.Migrate(gdb); err != nil {
		log.Fatalf("migrate: %v", err)
	}
	if err := seedIfEmpty(gdb); err != nil {
		log.Fatalf("seed: %v", err)
	}

	planRepo := mysqlrepo.NewPlanRepo(gdb)
	resvRepo := mysqlrepo.NewReservationRepo(gdb)
	userRepo := userrepo.NewUserRepo(gdb)

	reservationUC := &usecase.ReservationUsecase{Plans: planRepo, Resv: resvRepo, Users: userRepo}
	userUC := &usecase.UserUsecase{Users: userRepo}

	reservationHandler := &httpi.ReservationHandler{UC: reservationUC}
	userHandler := &httpi.UserHandler{UC: userUC}

	mux := http.NewServeMux()

	// 予約登録、予約一覧、予約取得、プラン検索、ユーザ登録API
	mux.HandleFunc("POST /reservations", reservationHandler.Create)
	mux.HandleFunc("GET /reservations", reservationHandler.List)
	mux.HandleFunc("GET /reservations/", reservationHandler.Get)
	mux.HandleFunc("GET /plans", reservationHandler.SearchPlans)
	mux.HandleFunc("POST /register", userHandler.Register)

	// ユーザ情報取得APIを追加
	mux.HandleFunc("GET /users/", userHandler.GetUser)

	addr := ":8080"
	log.Printf("listening on %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func seedIfEmpty(gdb *gorm.DB) error {
	var count int64
	if err := gdb.Model(&models.PlanModel{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	seed := []models.PlanModel{
		{ID: 100, Name: "富士プレミアム", Keyword: "富士 山 静岡", Price: 12000},
		{ID: 175, Name: "サウスベーシック", Keyword: "サウス 南", Price: 8000},
		{ID: 200, Name: "北の宿", Keyword: "北海道 北", Price: 10000},
	}
	return gdb.Create(&seed).Error
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
func getEnvInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return def
}
</file>

<file path="bookingapp/internal/domain/entity/plan.go">
package entity

type Plan struct {
	ID      int
	Name    string
	Keyword string // 簡易検索用
	Price   int
}
</file>

<file path="bookingapp/internal/domain/entity/reservation.go">
package entity

import "time"

type Reservation struct {
	ID       int
	UserID   string
	PlanID   int
	Number   int
	Checkin  time.Time
	Checkout time.Time
	Total    int // 計算済み合計金額
}

func (r *Reservation) Nights() int {
	d := r.Checkout.Sub(r.Checkin).Hours() / 24
	if d < 0 {
		return 0
	}
	return int(d)
}
</file>

<file path="bookingapp/internal/domain/entity/user.go">
package entity

import "time"

type User struct {
	ID           string    // ユーザーID
	Name         string    // 名前
	Email        string    // メールアドレス
	PhoneNumber  string    // 電話番号
	Address      string    // 住所（オプション）
	DateOfBirth  time.Time // 生年月日
	RegisteredAt time.Time // 登録日
	Status       string    // アカウントステータス（例: "active", "inactive"）
}
</file>

<file path="bookingapp/internal/domain/repository/repository.go">
package repository

import "bookingapp/internal/domain/entity"

type PlanRepository interface {
	FindByID(id int) (*entity.Plan, error)
	SearchByKeyword(keyword string) ([]*entity.Plan, error)
}

type ReservationRepository interface {
	NextID() int
	Save(reservation *entity.Reservation) (*entity.Reservation, error)
	FindByID(id int) (*entity.Reservation, error)
	List() ([]*entity.Reservation, error)
}

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Get(id string) (*entity.User, error)
}
</file>

<file path="bookingapp/internal/infrastructure/db/models/user/user_model.go">
package user

import "time"

type UserModel struct {
	ID           string     `gorm:"primaryKey;type:char(36)"`
	Name         string     `gorm:"size:255;not null"`
	Email        string     `gorm:"size:255;uniqueIndex;not null"`
	PhoneNumber  string     `gorm:"size:50"`
	Address      string     `gorm:"size:255"`
	DateOfBirth  *time.Time `gorm:"type:date"`
	RegisteredAt time.Time  `gorm:"not null"`
	Status       string     `gorm:"size:50;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (UserModel) TableName() string { return "users" }
</file>

<file path="bookingapp/internal/infrastructure/db/models/plan_model.go">
package models

import "time"

type PlanModel struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"size:255;not null"`
	Keyword   string `gorm:"size:255;index"`
	Price     int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (PlanModel) TableName() string { return "plans" }
</file>

<file path="bookingapp/internal/infrastructure/db/models/reservation_model.go">
package models

import "time"

type ReservationModel struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	UserID    string    `gorm:"type:char(36);not null;index"`
	PlanID    int       `gorm:"not null;index"`
	Number    int       `gorm:"not null"`
	Checkin   time.Time `gorm:"type:date;not null"`
	Checkout  time.Time `gorm:"type:date;not null"`
	Total     int       `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ReservationModel) TableName() string { return "reservations" }
</file>

<file path="bookingapp/internal/infrastructure/db/migrate.go">
package db

import (
	"bookingapp/internal/infrastructure/db/models"
	"bookingapp/internal/infrastructure/db/models/user" // UserModel をインポート
)

func Migrate(db any) error {
	gdb := db.(interface{ AutoMigrate(...any) error })
	return gdb.AutoMigrate(
		&models.PlanModel{},
		&models.ReservationModel{},
		&user.UserModel{}, // UserModel を追加
	)
}
</file>

<file path="bookingapp/internal/infrastructure/db/mysql.go">
package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	User string
	Pass string
	Host string // e.g. 127.0.0.1
	Port int    // e.g. 3306
	Name string // database name
}

func Open(c Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		c.User, c.Pass, c.Host, c.Port, c.Name,
	)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		// ここで必要ならLoggerやNamingStrategyを設定
	})
}

// 便利: Ping（接続確認）
func Ping(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	return sqlDB.Ping()
}
</file>

<file path="bookingapp/internal/infrastructure/memory/plan_repo_memory.go">
package memory

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	"strings"
)

type PlanRepoMemory struct {
	data map[int]*entity.Plan
}

func NewPlanRepoMemory(seed []*entity.Plan) repository.PlanRepository {
	m := &PlanRepoMemory{data: map[int]*entity.Plan{}}
	for _, p := range seed {
		cp := *p
		m.data[p.ID] = &cp
	}
	return m
}

func (m *PlanRepoMemory) FindByID(id int) (*entity.Plan, error) {
	if p, ok := m.data[id]; ok {
		cp := *p
		return &cp, nil
	}
	return nil, nil
}

func (m *PlanRepoMemory) SearchByKeyword(keyword string) ([]*entity.Plan, error) {
	if keyword == "" {
		out := make([]*entity.Plan, 0, len(m.data))
		for _, p := range m.data {
			cp := *p
			out = append(out, &cp)
		}
		return out, nil
	}
	kw := strings.ToLower(keyword)
	var out []*entity.Plan
	for _, p := range m.data {
		if strings.Contains(strings.ToLower(p.Name), kw) || strings.Contains(strings.ToLower(p.Keyword), kw) {
			cp := *p
			out = append(out, &cp)
		}
	}
	return out, nil
}
</file>

<file path="bookingapp/internal/infrastructure/memory/reservation_repo_memory.go">
package memory

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	"sync"
)

type ReservationRepoMemory struct {
	mu   sync.RWMutex
	data map[int]*entity.Reservation
	next int
}

func NewReservationRepoMemory() repository.ReservationRepository {
	return &ReservationRepoMemory{
		data: make(map[int]*entity.Reservation),
		next: 1,
	}
}

func (r *ReservationRepoMemory) NextID() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := r.next
	r.next++
	return id
}

func (r *ReservationRepoMemory) Save(res *entity.Reservation) (*entity.Reservation, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *res
	r.data[cp.ID] = &cp
	out := cp
	return &out, nil
}

func (r *ReservationRepoMemory) FindByID(id int) (*entity.Reservation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if v, ok := r.data[id]; ok {
		cp := *v
		return &cp, nil
	}
	return nil, nil
}

func (r *ReservationRepoMemory) List() ([]*entity.Reservation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*entity.Reservation, 0, len(r.data))
	for _, v := range r.data {
		cp := *v
		out = append(out, &cp)
	}
	return out, nil
}
</file>

<file path="bookingapp/internal/infrastructure/repository/mysql/user/user_get.go">
package user

import (
	"bookingapp/internal/domain/entity"
	usermodel "bookingapp/internal/infrastructure/db/models/user"
	"context"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// ---- ユーザー情報取得 ----
func (r *UserRepo) Get(id string) (*entity.User, error) {
	if strings.TrimSpace(id) == "" {
		return nil, fmt.Errorf("id is empty")
	}
	var model usermodel.UserModel
	err := r.db.WithContext(context.Background()).
		Where("id = ?", strings.TrimSpace(id)).
		First(&model).Error

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	case err != nil:
		return nil, err
	}

	return modelToEntity(&model), nil
}
</file>

<file path="bookingapp/internal/infrastructure/repository/mysql/user/user_repo_mysql.go">
package user

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	usermodel "bookingapp/internal/infrastructure/db/models/user"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// データベース操作を行うための実装
type UserRepo struct {
	db *gorm.DB
}

// 依存注入のためのコンストラクタ
func NewUserRepo(db *gorm.DB) repository.UserRepository {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *entity.User) (*entity.User, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}

	if user.ID == "" {
		user.ID = uuid.NewString()
	}

	var dob *time.Time
	if !user.DateOfBirth.IsZero() {
		d := user.DateOfBirth
		dob = &d
	}

	model := usermodel.UserModel{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		Address:      user.Address,
		DateOfBirth:  dob,
		RegisteredAt: user.RegisteredAt,
		Status:       user.Status,
	}

	if err := r.db.WithContext(context.Background()).Create(&model).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) FindByEmail(email string) (*entity.User, error) {
	if strings.TrimSpace(email) == "" {
		return nil, nil
	}

	var model usermodel.UserModel
	err := r.db.WithContext(context.Background()).
		Where("email = ?", strings.TrimSpace(email)).
		First(&model).Error

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	case err != nil:
		return nil, err
	}

	return modelToEntity(&model), nil
}

var _ repository.UserRepository = (*UserRepo)(nil)

func modelToEntity(model *usermodel.UserModel) *entity.User {
	if model == nil {
		return nil
	}

	var dob time.Time
	if model.DateOfBirth != nil {
		dob = *model.DateOfBirth
	}

	return &entity.User{
		ID:           model.ID,
		Name:         model.Name,
		Email:        model.Email,
		PhoneNumber:  model.PhoneNumber,
		Address:      model.Address,
		DateOfBirth:  dob,
		RegisteredAt: model.RegisteredAt,
		Status:       model.Status,
	}
}
</file>

<file path="bookingapp/internal/infrastructure/repository/mysql/plan_repo_mysql.go">
package mysqlrepo

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	"bookingapp/internal/infrastructure/db/models"
	"context"
	"strings"

	"gorm.io/gorm"
)

type PlanRepo struct{ db *gorm.DB }

func NewPlanRepo(db *gorm.DB) repository.PlanRepository { return &PlanRepo{db: db} }

func (r *PlanRepo) FindByID(id int) (*entity.Plan, error) {
	var m models.PlanModel
	if err := r.db.WithContext(context.Background()).
		First(&m, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &entity.Plan{ID: m.ID, Name: m.Name, Keyword: m.Keyword, Price: m.Price}, nil
}

func (r *PlanRepo) SearchByKeyword(keyword string) ([]*entity.Plan, error) {
	ctx := context.Background()
	var list []models.PlanModel

	q := r.db.WithContext(ctx).Model(&models.PlanModel{})
	if strings.TrimSpace(keyword) != "" {
		kw := "%" + strings.TrimSpace(keyword) + "%"
		q = q.Where("LOWER(name) LIKE LOWER(?) OR LOWER(keyword) LIKE LOWER(?)", kw, kw)
	}
	if err := q.Order("id ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*entity.Plan, 0, len(list))
	for _, m := range list {
		copy := m
		out = append(out, &entity.Plan{ID: copy.ID, Name: copy.Name, Keyword: copy.Keyword, Price: copy.Price})
	}
	return out, nil
}
</file>

<file path="bookingapp/internal/infrastructure/repository/mysql/reservation_repo_mysql.go">
package mysqlrepo

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	"bookingapp/internal/infrastructure/db/models"
	"context"

	"gorm.io/gorm"
)

type ReservationRepo struct{ db *gorm.DB }

func NewReservationRepo(db *gorm.DB) repository.ReservationRepository {
	return &ReservationRepo{db: db}
}

// DBのauto-incrementに委譲するのでNextIDは使わないが、interface満たすために実装
func (r *ReservationRepo) NextID() int { return 0 }

func (r *ReservationRepo) Save(res *entity.Reservation) (*entity.Reservation, error) {
	ctx := context.Background()
	m := models.ReservationModel{
		ID:       res.ID, // 0ならAUTO_INCREMENT
		UserID:   res.UserID,
		PlanID:   res.PlanID,
		Number:   res.Number,
		Checkin:  res.Checkin,
		Checkout: res.Checkout,
		Total:    res.Total,
	}
	if err := r.db.WithContext(ctx).Save(&m).Error; err != nil {
		return nil, err
	}
	// 生成されたIDを反映
	res.ID = m.ID
	return res, nil
}

func (r *ReservationRepo) FindByID(id int) (*entity.Reservation, error) {
	var m models.ReservationModel
	if err := r.db.WithContext(context.Background()).
		First(&m, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &entity.Reservation{
		ID:       m.ID,
		UserID:   m.UserID,
		PlanID:   m.PlanID,
		Number:   m.Number,
		Checkin:  m.Checkin,
		Checkout: m.Checkout,
		Total:    m.Total,
	}, nil
}

func (r *ReservationRepo) List() ([]*entity.Reservation, error) {
	var list []models.ReservationModel
	if err := r.db.WithContext(context.Background()).
		Order("id ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*entity.Reservation, 0, len(list))
	for _, m := range list {
		copy := m
		out = append(out, &entity.Reservation{
			ID:       copy.ID,
			UserID:   copy.UserID,
			PlanID:   copy.PlanID,
			Number:   copy.Number,
			Checkin:  copy.Checkin,
			Checkout: copy.Checkout,
			Total:    copy.Total,
		})
	}
	return out, nil
}

var _ repository.ReservationRepository = (*ReservationRepo)(nil)
</file>

<file path="bookingapp/internal/interface/http/reservation_handler.go">
package httpi

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/usecase"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ReservationHandler struct {
	UC *usecase.ReservationUsecase
}

type createReq struct {
	UserID   string `json:"user_id"`
	PlanID   int    `json:"plan_id"`
	Number   int    `json:"number"`
	Checkin  string `json:"checkin"`  // "2025-10-12"
	Checkout string `json:"checkout"` // "2025-10-13"
}

type createResp struct {
	ID int `json:"id"`
}

type reservationView struct {
	ID       int    `json:"id"`
	UserID   string `json:"user_id"`
	PlanID   int    `json:"plan_id"`
	Number   int    `json:"number"`
	Checkin  string `json:"checkin"`
	Checkout string `json:"checkout"`
	Total    int    `json:"total"`
	Nights   int    `json:"nights"`
}

func (h *ReservationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var in createReq
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	ci, err1 := time.Parse("2006-01-02", in.Checkin)
	co, err2 := time.Parse("2006-01-02", in.Checkout)
	if err1 != nil || err2 != nil {
		http.Error(w, "invalid date format (yyyy-mm-dd)", http.StatusBadRequest)
		return
	}
	res, err := h.UC.Create(in.UserID, in.PlanID, in.Number, ci, co)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrInvalidUserID):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, usecase.ErrUserNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)
		case errors.Is(err, usecase.ErrInvalidDates):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, usecase.ErrInvalidNumber):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, usecase.ErrPlanNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}
	writeJSON(w, http.StatusOK, createResp{ID: res.ID})
}

func (h *ReservationHandler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/reservations/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	res, _ := h.UC.Get(id)
	if res == nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, http.StatusOK, toView(res))
}

func (h *ReservationHandler) List(w http.ResponseWriter, r *http.Request) {
	list, _ := h.UC.List()
	views := make([]reservationView, 0, len(list))
	for _, v := range list {
		views = append(views, toView(v))
	}
	writeJSON(w, http.StatusOK, views)
}

func (h *ReservationHandler) SearchPlans(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("keyword")
	plans, _ := h.UC.SearchPlans(q)
	type planView struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Keyword string `json:"keyword"`
		Price   int    `json:"price"`
	}
	out := make([]planView, 0, len(plans))
	for _, p := range plans {
		out = append(out, planView{ID: p.ID, Name: p.Name, Keyword: p.Keyword, Price: p.Price})
	}
	writeJSON(w, http.StatusOK, out)
}

// ここを *entity.Reservation にする（別型を作らない）
func toView(r *entity.Reservation) reservationView {
	return reservationView{
		ID:       r.ID,
		UserID:   r.UserID,
		PlanID:   r.PlanID,
		Number:   r.Number,
		Checkin:  r.Checkin.Format("2006-01-02"),
		Checkout: r.Checkout.Format("2006-01-02"),
		Total:    r.Total,
		Nights:   r.Nights(), // entity に既にあるメソッドを使う
	}
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
</file>

<file path="bookingapp/internal/interface/http/user.go">
package httpi

import (
	"bookingapp/internal/usecase"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

type UserHandler struct {
	UC *usecase.UserUsecase
}

type registerUserReq struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
}

type registerUserResp struct {
	ID string `json:"id"`
}

type userView struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	Address      string `json:"address"`
	DateOfBirth  string `json:"date_of_birth"`
	RegisteredAt string `json:"registered_at"`
	Status       string `json:"status"`
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if h == nil || h.UC == nil {
		http.Error(w, "user lookup unavailable", http.StatusServiceUnavailable)
		return
	}

	id := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/users/"))
	if id == "" {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.UC.GetUser(id)
	if err != nil {
		if errors.Is(err, usecase.ErrUserInvalidInput) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.NotFound(w, r)
		return
	}

	view := userView{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		Address:      user.Address,
		DateOfBirth:  formatDate(user.DateOfBirth),
		RegisteredAt: user.RegisteredAt.Format(time.RFC3339),
		Status:       user.Status,
	}

	writeJSON(w, http.StatusOK, view)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if h == nil || h.UC == nil {
		http.Error(w, "user registration unavailable", http.StatusServiceUnavailable)
		return
	}

	var in registerUserReq
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	user, err := h.UC.Register(usecase.RegisterUserInput{
		Name:        in.Name,
		Email:       in.Email,
		PhoneNumber: in.PhoneNumber,
		Address:     in.Address,
		DateOfBirth: in.DateOfBirth,
	})
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrUserInvalidInput):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, usecase.ErrUserEmailAlreadyExists):
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	writeJSON(w, http.StatusCreated, registerUserResp{ID: user.ID})
}

func formatDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02")
}
</file>

<file path="bookingapp/internal/usecase/user/user_uc.go">
package usecase

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
)

type UserUsecase struct {
	Users repository.UserRepository
}

// ユーザー情報取得
func (u *UserUsecase) GetUser(id string) (*entity.User, error) {
	return u.Users.Get(id)
}
</file>

<file path="bookingapp/internal/usecase/reservation_uc.go">
package usecase

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidDates  = errors.New("invalid dates: checkout must be after checkin")
	ErrPlanNotFound  = errors.New("plan not found")
	ErrInvalidNumber = errors.New("number must be >= 1")
	ErrInvalidUserID = errors.New("invalid user id")
	ErrUserNotFound  = errors.New("user not found")
)

type ReservationUsecase struct {
	Users repository.UserRepository
	Plans repository.PlanRepository
	Resv  repository.ReservationRepository
}

// 　予約作成
func (u *ReservationUsecase) Create(userID string, planID, number int, checkin, checkout time.Time) (*entity.Reservation, error) {
	userID = strings.TrimSpace(userID)
	if userID == "" {
		return nil, ErrInvalidUserID
	}
	if _, err := uuid.Parse(userID); err != nil {
		return nil, ErrInvalidUserID
	}
	user, err := u.Users.Get(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	if !checkout.After(checkin) {
		return nil, ErrInvalidDates
	}
	if number < 1 {
		return nil, ErrInvalidNumber
	}
	plan, err := u.Plans.FindByID(planID)
	if err != nil {
		return nil, err
	}
	if plan == nil {
		return nil, ErrPlanNotFound
	}
	r := &entity.Reservation{
		ID:       u.Resv.NextID(),
		UserID:   user.ID,
		PlanID:   planID,
		Number:   number,
		Checkin:  checkin,
		Checkout: checkout,
	}
	//ドメイン層のメソッドを使って宿泊数を計算
	nights := r.Nights()
	//合計金額を計算してセット
	r.Total = plan.Price * number * nights
	//保存してID付きの予約情報を返す
	return u.Resv.Save(r)
}

// 予約取得
func (u *ReservationUsecase) Get(id int) (*entity.Reservation, error) {
	return u.Resv.FindByID(id)
}

// 予約一覧取得
func (u *ReservationUsecase) List() ([]*entity.Reservation, error) {
	return u.Resv.List()
}

// プラン検索
func (u *ReservationUsecase) SearchPlans(keyword string) ([]*entity.Plan, error) {
	return u.Plans.SearchByKeyword(keyword)
}
</file>

<file path="bookingapp/internal/usecase/user_uc.go">
package usecase

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	"errors"
	"strings"
	"time"
)

var (
	ErrUserInvalidInput       = errors.New("invalid user input")
	ErrUserEmailAlreadyExists = errors.New("user email already exists")
)

type RegisterUserInput struct {
	Name        string
	Email       string
	PhoneNumber string
	Address     string
	DateOfBirth string
}

// ユースケース層からrepository層のinterfaceを使えるようにする
type UserUsecase struct {
	Users repository.UserRepository
	Now   func() time.Time
}

func (u *UserUsecase) Register(in RegisterUserInput) (*entity.User, error) {
	if u.Users == nil {
		return nil, errors.New("user repository is nil")
	}

	name := strings.TrimSpace(in.Name)
	email := strings.TrimSpace(in.Email)
	if name == "" || email == "" {
		return nil, ErrUserInvalidInput
	}

	existing, err := u.Users.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, ErrUserEmailAlreadyExists
	}

	var dob time.Time
	if v := strings.TrimSpace(in.DateOfBirth); v != "" {
		dob, err = time.Parse("2006-01-02", v)
		if err != nil {
			return nil, ErrUserInvalidInput
		}
	}

	now := time.Now
	if u.Now != nil {
		now = u.Now
	}

	user := &entity.User{
		Name:         name,
		Email:        email,
		PhoneNumber:  strings.TrimSpace(in.PhoneNumber),
		Address:      strings.TrimSpace(in.Address),
		DateOfBirth:  dob,
		RegisteredAt: now(),
		Status:       "active",
	}

	return u.Users.Create(user)
}

func (u *UserUsecase) GetUser(id string) (*entity.User, error) {
	if u.Users == nil {
		return nil, errors.New("user repository is nil")
	}

	trimmed := strings.TrimSpace(id)
	if trimmed == "" {
		return nil, ErrUserInvalidInput
	}

	return u.Users.Get(trimmed)
}
</file>

<file path="bookingapp/docker-compose.yml">
version: "3.8"

services:
  mysql:
    image: mysql:8.0
    container_name: booking-mysql
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: booking
      MYSQL_USER: booking
      MYSQL_PASSWORD: bookingpass
      TZ: Asia/Tokyo
    ports:
      - "3306:3306"
    command:
      [
        "mysqld",
        "--default-authentication-plugin=mysql_native_password",
        "--character-set-server=utf8mb4",
        "--collation-server=utf8mb4_0900_ai_ci"
      ]
    volumes:
      - mysql_data:/var/lib/mysql
    networks:
      - booking-net

  api:
    build: .
    container_name: booking-api
    environment:
      DB_HOST: mysql
      DB_PORT: "3306"
      DB_USER: booking
      DB_PASS: bookingpass
      DB_NAME: booking
      PORT: "8080"
    depends_on:
      - mysql
    ports:
      - "8080:8080"
    networks:
      - booking-net

volumes:
  mysql_data: {}

networks:
  booking-net:
    driver: bridge
</file>

<file path="bookingapp/go.mod">
module bookingapp

go 1.24.0

require (
	gorm.io/driver/mysql v1.6.0
	gorm.io/gorm v1.31.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.3 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.30.0 // indirect
)
</file>

<file path="bookingapp/mermaid.md">
erDiagram
    USERS {
      char36 id PK "users.id"
      varchar name
      varchar email UK "unique"
      varchar phone_number
      varchar address
      date date_of_birth
      datetime registered_at
      varchar status
      datetime created_at
      datetime updated_at
    }

    PLANS {
      int id PK
      varchar name
      varchar keyword
      int price
      datetime created_at
      datetime updated_at
    }

    RESERVATIONS {
      int id PK "reservations.id"
      char36 user_id FK "-> users.id"
      int plan_id FK "-> plans.id"
      int number
      date checkin
      date checkout
      int total
      datetime created_at
      datetime updated_at
    }

    USERS ||--o{ RESERVATIONS : "users.id = reservations.user_id"
    PLANS ||--o{ RESERVATIONS : "plans.id = reservations.plan_id"
</file>

<file path="bookingapp/README.md">
# Booking App API

Go 言語と GORM を用いて宿泊予約を管理するシンプルな API サービスです。Clean Architecture を意識したレイヤ分割を採用し、ドメインロジックとインフラ依存コードを分離しています。

## 主な特徴
- MySQL（またはメモリリポジトリ）を利用したプラン・予約データ管理
- 予約作成時のバリデーション（宿泊日、人数、プラン存在チェック）
- RESTful なエンドポイント（プラン検索、予約 CRUD の一部）
- サーバ起動時の自動マイグレーションと初期データ投入

## ディレクトリ構成
```
.
├── cmd/api            # エントリーポイント（HTTP サーバ）
├── internal
│   ├── domain         # ドメインエンティティ & リポジトリインターフェース
│   ├── usecase        # ユースケース（アプリケーションサービス）
│   ├── interface/http # HTTP ハンドラ層
│   └── infrastructure # DB 接続・リポジトリ実装（MySQL / メモリ）
├── go.mod, go.sum
└── docker-compose.yml # MySQL 起動用定義
```

## アーキテクチャ概要
- `internal/domain/entity` にプラン (`Plan`) と予約 (`Reservation`) のドメインモデルを定義。`Reservation.Nights()` などのビジネスロジックをエンティティに寄せています。
- `internal/domain/repository` ではユースケースが依存するポート（インターフェース）を宣言。
- `internal/usecase/reservation_uc.go` はプラン検索や予約作成のアプリケーションロジックを担当し、入力バリデーションと料金計算を行います。
- `internal/interface/http` が HTTP リクエストを受け、ユースケースを呼び出して JSON を返却します。
- `internal/infrastructure` で具体的なアダプタを実装。`repository/mysql` は GORM を利用した永続化、`memory` はインメモリ実装です。

## 依存関係
- Go 1.24 以上
- MySQL 8 系（Docker Compose で起動可）
- ライブラリ
  - `gorm.io/gorm`
  - `gorm.io/driver/mysql`

## 環境変数
`cmd/api/main.go` では以下の環境変数を読み込み、未設定の場合はデフォルト値を使用します。

| 変数名   | デフォルト値  | 説明                 |
|----------|----------------|----------------------|
| `DB_HOST`| `127.0.0.1`    | MySQL ホスト         |
| `DB_PORT`| `3306`         | MySQL ポート         |
| `DB_USER`| `root`         | 接続ユーザー         |
| `DB_PASS`| `password`     | 接続パスワード       |
| `DB_NAME`| `booking`      | 使用するデータベース |

## 起動方法
1. 依存環境を用意
   ```bash
   docker compose up -d mysql
   ```
   > `docker-compose.yml` は root パスワード `password`、DB 名 `booking` で MySQL 8.0 を立ち上げます。

2. API サーバを起動
   ```bash
   go run ./cmd/api
   ```
   起動すると `:8080` で HTTP サーバが待ち受けます。

### マイグレーションとシード
サーバ起動時に以下が自動で実行されます。
- GORM の `AutoMigrate` による `plans` / `reservations` テーブル生成
- `plans` テーブルが空の場合、初期プラン 3 件を投入
  - 例: `ID=100, Name="富士プレミアム", Price=12000`

## ドメインロジック
- 予約作成 (`ReservationUsecase.Create`)
  - チェックイン < チェックアウト、人数 >= 1 を検証
  - 指定プランの存在確認（未存在時は `ErrPlanNotFound`）
  - 宿泊数（`Reservation.Nights()`）と人数、プラン単価から合計金額を算出
  - リポジトリ経由で保存し、生成された ID を返却
- 予約参照 (`Get`, `List`) とプラン検索 (`SearchPlans`) もユースケースを経由

## HTTP API
| メソッド | パス                | 説明                           |
|----------|---------------------|--------------------------------|
| `POST`   | `/reservations`     | 予約を新規作成                 |
| `GET`    | `/reservations`     | 予約一覧を取得                 |
| `GET`    | `/reservations/{id}`| 予約詳細を取得                 |
| `GET`    | `/plans`            | キーワードでプランを検索       |

### リクエスト/レスポンス例
**予約作成**
```bash
curl -X POST http://localhost:8080/reservations \
  -H 'Content-Type: application/json' \
  -d '{
        "plan_id": 100,
        "number": 2,
        "checkin": "2025-10-12",
        "checkout": "2025-10-14"
      }'
```
レスポンス
```json
{ "id": 1 }
```

**予約一覧**
```bash
curl http://localhost:8080/reservations
```
レスポンス（例）
```json
[
  {
    "id": 1,
    "plan_id": 100,
    "number": 2,
    "checkin": "2025-10-12",
    "checkout": "2025-10-14",
    "total": 48000,
    "nights": 2
  }
]
```

**プラン検索**
```bash
curl "http://localhost:8080/plans?keyword=富士"
```
レスポンス（例）
```json
[
  {
    "id": 100,
    "name": "富士プレミアム",
    "keyword": "富士 山 静岡",
    "price": 12000
  }
]
```

### エラーレスポンス
- リクエスト JSON / 日付フォーマット不正: `400 Bad Request`
- 宿泊日逆転・人数不足: `400 Bad Request`
- 存在しないプラン指定: `404 Not Found`
- その他予期しないエラー: `500 Internal Server Error`

## テストや拡張のヒント
- インメモリリポジトリ（`internal/infrastructure/memory`）を利用してユニットテストを書けます。
- バリデーション強化（例: 最大人数、予約重複チェック）や、キャンセル API 追加などの拡張が容易です。
- HTTP レイヤは `net/http` 標準ライブラリのままなので、Echo や Chi などに置き換える場合もユースケース層は流用可能です。

## ライセンス
本リポジトリにライセンス表記がない場合は、利用前に作成者へ確認してください。
</file>

<file path="infra/bin/infra.ts">
#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib';
import { InfraStack } from '../lib/infra-stack';

const app = new cdk.App();
new InfraStack(app, 'InfraStack', {
  /* If you don't specify 'env', this stack will be environment-agnostic.
   * Account/Region-dependent features and context lookups will not work,
   * but a single synthesized template can be deployed anywhere. */

  /* Uncomment the next line to specialize this stack for the AWS Account
   * and Region that are implied by the current CLI configuration. */
  // env: { account: process.env.CDK_DEFAULT_ACCOUNT, region: process.env.CDK_DEFAULT_REGION },

  /* Uncomment the next line if you know exactly what Account and Region you
   * want to deploy the stack to. */
  // env: { account: '123456789012', region: 'us-east-1' },

  /* For more information, see https://docs.aws.amazon.com/cdk/latest/guide/environments.html */
});
</file>

<file path="infra/lib/infra-stack.ts">
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
</file>

<file path="infra/test/infra.test.ts">
// import * as cdk from 'aws-cdk-lib';
// import { Template } from 'aws-cdk-lib/assertions';
// import * as Infra from '../lib/infra-stack';

// example test. To run these tests, uncomment this file along with the
// example resource in lib/infra-stack.ts
test('SQS Queue Created', () => {
//   const app = new cdk.App();
//     // WHEN
//   const stack = new Infra.InfraStack(app, 'MyTestStack');
//     // THEN
//   const template = Template.fromStack(stack);

//   template.hasResourceProperties('AWS::SQS::Queue', {
//     VisibilityTimeout: 300
//   });
});
</file>

<file path="infra/.gitignore">
*.js
!jest.config.js
*.d.ts
node_modules

# CDK asset staging directory
.cdk.staging
cdk.out
</file>

<file path="infra/.npmignore">
*.ts
!*.d.ts

# CDK asset staging directory
.cdk.staging
cdk.out
</file>

<file path="infra/cdk.json">
{
  "app": "npx ts-node --prefer-ts-exts bin/infra.ts",
  "watch": {
    "include": [
      "**"
    ],
    "exclude": [
      "README.md",
      "cdk*.json",
      "**/*.d.ts",
      "**/*.js",
      "tsconfig.json",
      "package*.json",
      "yarn.lock",
      "node_modules",
      "test"
    ]
  },
  "context": {
    "@aws-cdk/aws-signer:signingProfileNamePassedToCfn": true,
    "@aws-cdk/aws-ecs-patterns:secGroupsDisablesImplicitOpenListener": true,
    "@aws-cdk/aws-lambda:recognizeLayerVersion": true,
    "@aws-cdk/core:checkSecretUsage": true,
    "@aws-cdk/core:target-partitions": [
      "aws",
      "aws-cn"
    ],
    "@aws-cdk-containers/ecs-service-extensions:enableDefaultLogDriver": true,
    "@aws-cdk/aws-ec2:uniqueImdsv2TemplateName": true,
    "@aws-cdk/aws-ecs:arnFormatIncludesClusterName": true,
    "@aws-cdk/aws-iam:minimizePolicies": true,
    "@aws-cdk/core:validateSnapshotRemovalPolicy": true,
    "@aws-cdk/aws-codepipeline:crossAccountKeyAliasStackSafeResourceName": true,
    "@aws-cdk/aws-s3:createDefaultLoggingPolicy": true,
    "@aws-cdk/aws-sns-subscriptions:restrictSqsDescryption": true,
    "@aws-cdk/aws-apigateway:disableCloudWatchRole": true,
    "@aws-cdk/core:enablePartitionLiterals": true,
    "@aws-cdk/aws-events:eventsTargetQueueSameAccount": true,
    "@aws-cdk/aws-ecs:disableExplicitDeploymentControllerForCircuitBreaker": true,
    "@aws-cdk/aws-iam:importedRoleStackSafeDefaultPolicyName": true,
    "@aws-cdk/aws-s3:serverAccessLogsUseBucketPolicy": true,
    "@aws-cdk/aws-route53-patters:useCertificate": true,
    "@aws-cdk/customresources:installLatestAwsSdkDefault": false,
    "@aws-cdk/aws-rds:databaseProxyUniqueResourceName": true,
    "@aws-cdk/aws-codedeploy:removeAlarmsFromDeploymentGroup": true,
    "@aws-cdk/aws-apigateway:authorizerChangeDeploymentLogicalId": true,
    "@aws-cdk/aws-ec2:launchTemplateDefaultUserData": true,
    "@aws-cdk/aws-secretsmanager:useAttachedSecretResourcePolicyForSecretTargetAttachments": true,
    "@aws-cdk/aws-redshift:columnId": true,
    "@aws-cdk/aws-stepfunctions-tasks:enableEmrServicePolicyV2": true,
    "@aws-cdk/aws-ec2:restrictDefaultSecurityGroup": true,
    "@aws-cdk/aws-apigateway:requestValidatorUniqueId": true,
    "@aws-cdk/aws-kms:aliasNameRef": true,
    "@aws-cdk/aws-kms:applyImportedAliasPermissionsToPrincipal": true,
    "@aws-cdk/aws-autoscaling:generateLaunchTemplateInsteadOfLaunchConfig": true,
    "@aws-cdk/core:includePrefixInUniqueNameGeneration": true,
    "@aws-cdk/aws-efs:denyAnonymousAccess": true,
    "@aws-cdk/aws-opensearchservice:enableOpensearchMultiAzWithStandby": true,
    "@aws-cdk/aws-lambda-nodejs:useLatestRuntimeVersion": true,
    "@aws-cdk/aws-efs:mountTargetOrderInsensitiveLogicalId": true,
    "@aws-cdk/aws-rds:auroraClusterChangeScopeOfInstanceParameterGroupWithEachParameters": true,
    "@aws-cdk/aws-appsync:useArnForSourceApiAssociationIdentifier": true,
    "@aws-cdk/aws-rds:preventRenderingDeprecatedCredentials": true,
    "@aws-cdk/aws-codepipeline-actions:useNewDefaultBranchForCodeCommitSource": true,
    "@aws-cdk/aws-cloudwatch-actions:changeLambdaPermissionLogicalIdForLambdaAction": true,
    "@aws-cdk/aws-codepipeline:crossAccountKeysDefaultValueToFalse": true,
    "@aws-cdk/aws-codepipeline:defaultPipelineTypeToV2": true,
    "@aws-cdk/aws-kms:reduceCrossAccountRegionPolicyScope": true,
    "@aws-cdk/aws-eks:nodegroupNameAttribute": true,
    "@aws-cdk/aws-ec2:ebsDefaultGp3Volume": true,
    "@aws-cdk/aws-ecs:removeDefaultDeploymentAlarm": true,
    "@aws-cdk/custom-resources:logApiResponseDataPropertyTrueDefault": false,
    "@aws-cdk/aws-s3:keepNotificationInImportedBucket": false,
    "@aws-cdk/core:explicitStackTags": true,
    "@aws-cdk/aws-ecs:enableImdsBlockingDeprecatedFeature": false,
    "@aws-cdk/aws-ecs:disableEcsImdsBlocking": true,
    "@aws-cdk/aws-ecs:reduceEc2FargateCloudWatchPermissions": true,
    "@aws-cdk/aws-dynamodb:resourcePolicyPerReplica": true,
    "@aws-cdk/aws-ec2:ec2SumTImeoutEnabled": true,
    "@aws-cdk/aws-appsync:appSyncGraphQLAPIScopeLambdaPermission": true,
    "@aws-cdk/aws-rds:setCorrectValueForDatabaseInstanceReadReplicaInstanceResourceId": true,
    "@aws-cdk/core:cfnIncludeRejectComplexResourceUpdateCreatePolicyIntrinsics": true,
    "@aws-cdk/aws-lambda-nodejs:sdkV3ExcludeSmithyPackages": true,
    "@aws-cdk/aws-stepfunctions-tasks:fixRunEcsTaskPolicy": true,
    "@aws-cdk/aws-ec2:bastionHostUseAmazonLinux2023ByDefault": true,
    "@aws-cdk/aws-route53-targets:userPoolDomainNameMethodWithoutCustomResource": true,
    "@aws-cdk/aws-elasticloadbalancingV2:albDualstackWithoutPublicIpv4SecurityGroupRulesDefault": true,
    "@aws-cdk/aws-iam:oidcRejectUnauthorizedConnections": true,
    "@aws-cdk/core:enableAdditionalMetadataCollection": true,
    "@aws-cdk/aws-lambda:createNewPoliciesWithAddToRolePolicy": false,
    "@aws-cdk/aws-s3:setUniqueReplicationRoleName": true,
    "@aws-cdk/aws-events:requireEventBusPolicySid": true,
    "@aws-cdk/core:aspectPrioritiesMutating": true,
    "@aws-cdk/aws-dynamodb:retainTableReplica": true,
    "@aws-cdk/aws-stepfunctions:useDistributedMapResultWriterV2": true,
    "@aws-cdk/s3-notifications:addS3TrustKeyPolicyForSnsSubscriptions": true,
    "@aws-cdk/aws-ec2:requirePrivateSubnetsForEgressOnlyInternetGateway": true,
    "@aws-cdk/aws-s3:publicAccessBlockedByDefault": true,
    "@aws-cdk/aws-lambda:useCdkManagedLogGroup": true
  }
}
</file>

<file path="infra/jest.config.js">
module.exports = {
  testEnvironment: 'node',
  roots: ['<rootDir>/test'],
  testMatch: ['**/*.test.ts'],
  transform: {
    '^.+\\.tsx?$': 'ts-jest'
  }
};
</file>

<file path="infra/package.json">
{
  "name": "infra",
  "version": "0.1.0",
  "bin": {
    "infra": "bin/infra.js"
  },
  "scripts": {
    "build": "tsc",
    "watch": "tsc -w",
    "test": "jest",
    "cdk": "cdk"
  },
  "devDependencies": {
    "@types/jest": "^29.5.14",
    "@types/node": "22.7.9",
    "jest": "^29.7.0",
    "ts-jest": "^29.2.5",
    "aws-cdk": "2.1030.0",
    "ts-node": "^10.9.2",
    "typescript": "~5.6.3"
  },
  "dependencies": {
    "aws-cdk-lib": "2.215.0",
    "constructs": "^10.0.0"
  }
}
</file>

<file path="infra/README.md">
# Welcome to your CDK TypeScript project

This is a blank project for CDK development with TypeScript.

The `cdk.json` file tells the CDK Toolkit how to execute your app.

## Useful commands

* `npm run build`   compile typescript to js
* `npm run watch`   watch for changes and compile
* `npm run test`    perform the jest unit tests
* `npx cdk deploy`  deploy this stack to your default AWS account/region
* `npx cdk diff`    compare deployed stack with current state
* `npx cdk synth`   emits the synthesized CloudFormation template
</file>

<file path="infra/tsconfig.json">
{
  "compilerOptions": {
    "target": "ES2022",
    "module": "NodeNext",
    "moduleResolution": "NodeNext",
    "lib": [
      "es2022"
    ],
    "declaration": true,
    "strict": true,
    "noImplicitAny": true,
    "strictNullChecks": true,
    "noImplicitThis": true,
    "alwaysStrict": true,
    "noUnusedLocals": false,
    "noUnusedParameters": false,
    "noImplicitReturns": true,
    "noFallthroughCasesInSwitch": false,
    "inlineSourceMap": true,
    "inlineSources": true,
    "experimentalDecorators": true,
    "strictPropertyInitialization": false,
    "skipLibCheck": true,
    "typeRoots": [
      "./node_modules/@types"
    ]
  },
  "exclude": [
    "node_modules",
    "cdk.out"
  ]
}
</file>

<file path="fix.md">
ユーザー会員登録

 - bookingapp/internal/domain/repository/repository.go:17 に UserRepository を追加し、ユースケースが依存するインターフェースを定義。
  - bookingapp/internal/usecase/user_uc.go:29 で登録ロジックを実装し、入力バリデーション・日付パース・重複メール検出を行って UserRepository へ委譲。
  - bookingapp/internal/infrastructure/repository/mysql/user/user_repo_mysql.go:24 に MySQL 実装を追加し、UUID 発行・DateOfBirth の NULL 対応・既存メール検索を実装。
  - bookingapp/internal/interface/http/user.go:26 にユーザーハンドラを作成し、HTTP からユースケースを呼び出して適切な HTTP ステータスを返却。
  - bookingapp/cmd/api/main.go:42 で新しいハンドラ/ユースケース/リポジトリを初期化し、POST /register を UserHandler にバインド。
</file>

</files>
````

## File: .serena/memories/project_overview.md
````markdown
# BookingApp Overview
- Purpose: Clean Architecture sample for hotel reservation bookings with RESTful HTTP API backed by MySQL via GORM.
- Tech stack: Go 1.24+ modules, net/http, GORM (mysql driver), optional in-memory repos for tests.
- Structure: `cmd/api` bootstrap, `internal/domain` entities + repository contracts, `internal/usecase` orchestrates, `internal/interface/http` HTTP handlers, `internal/infrastructure/db` GORM models + migrations, `internal/infrastructure/memory` in-memory repos.
- Key patterns: layered Clean Architecture with repositories injected into usecases, HTTP handlers consuming usecases, optional DB seeding on startup.
- Running env: expects MySQL reachable via env vars `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASS`, `DB_NAME`; seeds sample plans if table empty.
````

## File: .serena/memories/style_and_conventions.md
````markdown
# Style and Conventions
- Standard Go style: gofmt formatting, short receiver names, exported types for domain entities/usecases, unexported struct fields where possible.
- Clean Architecture layering: domain entities free of infra deps; usecases depend on domain repository interfaces; interface/http packages convert transport payloads.
- Error handling: return Go errors, map to HTTP status codes in handlers; prefer sentinel errors (`ErrInvalidDates`, etc.).
- Repositories expose interfaces for Plan and Reservation storage; concrete adapters live under infrastructure (MySQL, in-memory).
- Testing not present yet; in-memory repos can support unit tests without DB.
- Comments primarily Japanese/English mix, keep brief purposeful notes when logic not obvious.
````

## File: .serena/memories/suggested_commands.md
````markdown
# Suggested Commands
- `go run ./cmd/api` – start the HTTP API (uses env vars `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASS`, `DB_NAME`; defaults to localhost MySQL, auto-migrates + seeds plans).
- `go build ./cmd/api` – compile the server binary.
- `go test ./...` – run all Go unit tests (none yet, but command stays standard).
- `gofmt -w <files>` – format Go source before committing.
- `golangci-lint run` – if linting is added; current project does not vendor config.
````

## File: .serena/memories/task_completion.md
````markdown
# Task Completion Checklist
- Ensure modified Go files are gofmt formatted (`gofmt -w`).
- Run `go test ./...` to confirm unit tests (once they exist) pass.
- If server behavior affected, consider `go run ./cmd/api` smoke test against local MySQL.
- Update documentation/comments when changing exposed behavior or contracts.
- Summarize changes and note any manual verification steps for the user.
````

## File: .serena/.gitignore
````
/cache
````

## File: .serena/project.yml
````yaml
# language of the project (csharp, python, rust, java, typescript, go, cpp, or ruby)
#  * For C, use cpp
#  * For JavaScript, use typescript
# Special requirements:
#  * csharp: Requires the presence of a .sln file in the project folder.
language: go

# whether to use the project's gitignore file to ignore files
# Added on 2025-04-07
ignore_all_files_in_gitignore: true
# list of additional paths to ignore
# same syntax as gitignore, so you can use * and **
# Was previously called `ignored_dirs`, please update your config if you are using that.
# Added (renamed) on 2025-04-07
ignored_paths: []

# whether the project is in read-only mode
# If set to true, all editing tools will be disabled and attempts to use them will result in an error
# Added on 2025-04-18
read_only: false

# list of tool names to exclude. We recommend not excluding any tools, see the readme for more details.
# Below is the complete list of tools for convenience.
# To make sure you have the latest list of tools, and to view their descriptions, 
# execute `uv run scripts/print_tool_overview.py`.
#
#  * `activate_project`: Activates a project by name.
#  * `check_onboarding_performed`: Checks whether project onboarding was already performed.
#  * `create_text_file`: Creates/overwrites a file in the project directory.
#  * `delete_lines`: Deletes a range of lines within a file.
#  * `delete_memory`: Deletes a memory from Serena's project-specific memory store.
#  * `execute_shell_command`: Executes a shell command.
#  * `find_referencing_code_snippets`: Finds code snippets in which the symbol at the given location is referenced.
#  * `find_referencing_symbols`: Finds symbols that reference the symbol at the given location (optionally filtered by type).
#  * `find_symbol`: Performs a global (or local) search for symbols with/containing a given name/substring (optionally filtered by type).
#  * `get_current_config`: Prints the current configuration of the agent, including the active and available projects, tools, contexts, and modes.
#  * `get_symbols_overview`: Gets an overview of the top-level symbols defined in a given file.
#  * `initial_instructions`: Gets the initial instructions for the current project.
#     Should only be used in settings where the system prompt cannot be set,
#     e.g. in clients you have no control over, like Claude Desktop.
#  * `insert_after_symbol`: Inserts content after the end of the definition of a given symbol.
#  * `insert_at_line`: Inserts content at a given line in a file.
#  * `insert_before_symbol`: Inserts content before the beginning of the definition of a given symbol.
#  * `list_dir`: Lists files and directories in the given directory (optionally with recursion).
#  * `list_memories`: Lists memories in Serena's project-specific memory store.
#  * `onboarding`: Performs onboarding (identifying the project structure and essential tasks, e.g. for testing or building).
#  * `prepare_for_new_conversation`: Provides instructions for preparing for a new conversation (in order to continue with the necessary context).
#  * `read_file`: Reads a file within the project directory.
#  * `read_memory`: Reads the memory with the given name from Serena's project-specific memory store.
#  * `remove_project`: Removes a project from the Serena configuration.
#  * `replace_lines`: Replaces a range of lines within a file with new content.
#  * `replace_symbol_body`: Replaces the full definition of a symbol.
#  * `restart_language_server`: Restarts the language server, may be necessary when edits not through Serena happen.
#  * `search_for_pattern`: Performs a search for a pattern in the project.
#  * `summarize_changes`: Provides instructions for summarizing the changes made to the codebase.
#  * `switch_modes`: Activates modes by providing a list of their names
#  * `think_about_collected_information`: Thinking tool for pondering the completeness of collected information.
#  * `think_about_task_adherence`: Thinking tool for determining whether the agent is still on track with the current task.
#  * `think_about_whether_you_are_done`: Thinking tool for determining whether the task is truly completed.
#  * `write_memory`: Writes a named memory (for future reference) to Serena's project-specific memory store.
excluded_tools: []

# initial prompt for the project. It will always be given to the LLM upon activating the project
# (contrary to the memories, which are loaded on demand).
initial_prompt: ""

project_name: "CleanArc"
````

## File: bookingapp/cmd/api/main.go
````go
package main

import (
	"bookingapp/internal/infrastructure/db"
	"bookingapp/internal/infrastructure/db/models"
	mysqlrepo "bookingapp/internal/infrastructure/repository/mysql"
	userrepo "bookingapp/internal/infrastructure/repository/mysql/user"
	httpi "bookingapp/internal/interface/http"
	"bookingapp/internal/usecase"
	"log"
	"net/http"
	"os"
	"strconv"

	"gorm.io/gorm"
)

func main() {
	// ---- 環境変数から接続情報 ----
	host := getEnv("DB_HOST", "127.0.0.1")
	port := getEnvInt("DB_PORT", 3306)
	user := getEnv("DB_USER", "root")
	pass := getEnv("DB_PASS", "password")
	name := getEnv("DB_NAME", "booking")

	gdb, err := db.Open(db.Config{
		User: user, Pass: pass, Host: host, Port: port, Name: name,
	})
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	if err := db.Ping(gdb); err != nil {
		log.Fatalf("ping db: %v", err)
	}
	if err := db.Migrate(gdb); err != nil {
		log.Fatalf("migrate: %v", err)
	}
	if err := seedIfEmpty(gdb); err != nil {
		log.Fatalf("seed: %v", err)
	}

	planRepo := mysqlrepo.NewPlanRepo(gdb)
	resvRepo := mysqlrepo.NewReservationRepo(gdb)
	userRepo := userrepo.NewUserRepo(gdb)

	reservationUC := &usecase.ReservationUsecase{Plans: planRepo, Resv: resvRepo, Users: userRepo}
	userUC := &usecase.UserUsecase{Users: userRepo}

	reservationHandler := &httpi.ReservationHandler{UC: reservationUC}
	userHandler := &httpi.UserHandler{UC: userUC}

	mux := http.NewServeMux()

	// 予約登録、予約一覧、予約取得、プラン検索、ユーザ登録API
	mux.HandleFunc("POST /reservations", reservationHandler.Create)
	mux.HandleFunc("GET /reservations", reservationHandler.List)
	mux.HandleFunc("GET /reservations/", reservationHandler.Get)
	mux.HandleFunc("GET /plans", reservationHandler.SearchPlans)
	mux.HandleFunc("POST /register", userHandler.Register)

	// ユーザ情報取得APIを追加
	mux.HandleFunc("GET /users/", userHandler.GetUser)

	addr := ":8080"
	log.Printf("listening on %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func seedIfEmpty(gdb *gorm.DB) error {
	var count int64
	if err := gdb.Model(&models.PlanModel{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	seed := []models.PlanModel{
		{ID: 100, Name: "富士プレミアム", Keyword: "富士 山 静岡", Price: 12000},
		{ID: 175, Name: "サウスベーシック", Keyword: "サウス 南", Price: 8000},
		{ID: 200, Name: "北の宿", Keyword: "北海道 北", Price: 10000},
	}
	return gdb.Create(&seed).Error
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
func getEnvInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return def
}
````

## File: bookingapp/internal/domain/entity/plan.go
````go
package entity

type Plan struct {
	ID      int
	Name    string
	Keyword string // 簡易検索用
	Price   int
}
````

## File: bookingapp/internal/domain/entity/reservation.go
````go
package entity

import "time"

type Reservation struct {
	ID       int
	UserID   string
	PlanID   int
	Number   int
	Checkin  time.Time
	Checkout time.Time
	Total    int // 計算済み合計金額
}

func (r *Reservation) Nights() int {
	d := r.Checkout.Sub(r.Checkin).Hours() / 24
	if d < 0 {
		return 0
	}
	return int(d)
}
````

## File: bookingapp/internal/domain/entity/user.go
````go
package entity

import "time"

type User struct {
	ID           string    // ユーザーID
	Name         string    // 名前
	Email        string    // メールアドレス
	PhoneNumber  string    // 電話番号
	Address      string    // 住所（オプション）
	DateOfBirth  time.Time // 生年月日
	RegisteredAt time.Time // 登録日
	Status       string    // アカウントステータス（例: "active", "inactive"）
}
````

## File: bookingapp/internal/domain/repository/repository.go
````go
package repository

import "bookingapp/internal/domain/entity"

type PlanRepository interface {
	FindByID(id int) (*entity.Plan, error)
	SearchByKeyword(keyword string) ([]*entity.Plan, error)
}

type ReservationRepository interface {
	NextID() int
	Save(reservation *entity.Reservation) (*entity.Reservation, error)
	FindByID(id int) (*entity.Reservation, error)
	List() ([]*entity.Reservation, error)
}

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Get(id string) (*entity.User, error)
}
````

## File: bookingapp/internal/infrastructure/db/models/user/user_model.go
````go
package user

import "time"

type UserModel struct {
	ID           string     `gorm:"primaryKey;type:char(36)"`
	Name         string     `gorm:"size:255;not null"`
	Email        string     `gorm:"size:255;uniqueIndex;not null"`
	PhoneNumber  string     `gorm:"size:50"`
	Address      string     `gorm:"size:255"`
	DateOfBirth  *time.Time `gorm:"type:date"`
	RegisteredAt time.Time  `gorm:"not null"`
	Status       string     `gorm:"size:50;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (UserModel) TableName() string { return "users" }
````

## File: bookingapp/internal/infrastructure/db/models/plan_model.go
````go
package models

import "time"

type PlanModel struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"size:255;not null"`
	Keyword   string `gorm:"size:255;index"`
	Price     int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (PlanModel) TableName() string { return "plans" }
````

## File: bookingapp/internal/infrastructure/db/models/reservation_model.go
````go
package models

import "time"

type ReservationModel struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	UserID    string    `gorm:"type:char(36);not null;index"`
	PlanID    int       `gorm:"not null;index"`
	Number    int       `gorm:"not null"`
	Checkin   time.Time `gorm:"type:date;not null"`
	Checkout  time.Time `gorm:"type:date;not null"`
	Total     int       `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ReservationModel) TableName() string { return "reservations" }
````

## File: bookingapp/internal/infrastructure/db/migrate.go
````go
package db

import (
	"bookingapp/internal/infrastructure/db/models"
	"bookingapp/internal/infrastructure/db/models/user" // UserModel をインポート
)

func Migrate(db any) error {
	gdb := db.(interface{ AutoMigrate(...any) error })
	return gdb.AutoMigrate(
		&models.PlanModel{},
		&models.ReservationModel{},
		&user.UserModel{}, // UserModel を追加
	)
}
````

## File: bookingapp/internal/infrastructure/db/mysql.go
````go
package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	User string
	Pass string
	Host string // e.g. 127.0.0.1
	Port int    // e.g. 3306
	Name string // database name
}

func Open(c Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		c.User, c.Pass, c.Host, c.Port, c.Name,
	)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		// ここで必要ならLoggerやNamingStrategyを設定
	})
}

// 便利: Ping（接続確認）
func Ping(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	return sqlDB.Ping()
}
````

## File: bookingapp/internal/infrastructure/memory/plan_repo_memory.go
````go
package memory

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	"strings"
)

type PlanRepoMemory struct {
	data map[int]*entity.Plan
}

func NewPlanRepoMemory(seed []*entity.Plan) repository.PlanRepository {
	m := &PlanRepoMemory{data: map[int]*entity.Plan{}}
	for _, p := range seed {
		cp := *p
		m.data[p.ID] = &cp
	}
	return m
}

func (m *PlanRepoMemory) FindByID(id int) (*entity.Plan, error) {
	if p, ok := m.data[id]; ok {
		cp := *p
		return &cp, nil
	}
	return nil, nil
}

func (m *PlanRepoMemory) SearchByKeyword(keyword string) ([]*entity.Plan, error) {
	if keyword == "" {
		out := make([]*entity.Plan, 0, len(m.data))
		for _, p := range m.data {
			cp := *p
			out = append(out, &cp)
		}
		return out, nil
	}
	kw := strings.ToLower(keyword)
	var out []*entity.Plan
	for _, p := range m.data {
		if strings.Contains(strings.ToLower(p.Name), kw) || strings.Contains(strings.ToLower(p.Keyword), kw) {
			cp := *p
			out = append(out, &cp)
		}
	}
	return out, nil
}
````

## File: bookingapp/internal/infrastructure/memory/reservation_repo_memory.go
````go
package memory

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	"sync"
)

type ReservationRepoMemory struct {
	mu   sync.RWMutex
	data map[int]*entity.Reservation
	next int
}

func NewReservationRepoMemory() repository.ReservationRepository {
	return &ReservationRepoMemory{
		data: make(map[int]*entity.Reservation),
		next: 1,
	}
}

func (r *ReservationRepoMemory) NextID() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := r.next
	r.next++
	return id
}

func (r *ReservationRepoMemory) Save(res *entity.Reservation) (*entity.Reservation, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *res
	r.data[cp.ID] = &cp
	out := cp
	return &out, nil
}

func (r *ReservationRepoMemory) FindByID(id int) (*entity.Reservation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if v, ok := r.data[id]; ok {
		cp := *v
		return &cp, nil
	}
	return nil, nil
}

func (r *ReservationRepoMemory) List() ([]*entity.Reservation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*entity.Reservation, 0, len(r.data))
	for _, v := range r.data {
		cp := *v
		out = append(out, &cp)
	}
	return out, nil
}
````

## File: bookingapp/internal/infrastructure/repository/mysql/user/user_get.go
````go
package user

import (
	"bookingapp/internal/domain/entity"
	usermodel "bookingapp/internal/infrastructure/db/models/user"
	"context"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// ---- ユーザー情報取得 ----
func (r *UserRepo) Get(id string) (*entity.User, error) {
	if strings.TrimSpace(id) == "" {
		return nil, fmt.Errorf("id is empty")
	}
	var model usermodel.UserModel
	err := r.db.WithContext(context.Background()).
		Where("id = ?", strings.TrimSpace(id)).
		First(&model).Error

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	case err != nil:
		return nil, err
	}

	return modelToEntity(&model), nil
}
````

## File: bookingapp/internal/infrastructure/repository/mysql/user/user_repo_mysql.go
````go
package user

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	usermodel "bookingapp/internal/infrastructure/db/models/user"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// データベース操作を行うための実装
type UserRepo struct {
	db *gorm.DB
}

// 依存注入のためのコンストラクタ
func NewUserRepo(db *gorm.DB) repository.UserRepository {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *entity.User) (*entity.User, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}

	if user.ID == "" {
		user.ID = uuid.NewString()
	}

	var dob *time.Time
	if !user.DateOfBirth.IsZero() {
		d := user.DateOfBirth
		dob = &d
	}

	model := usermodel.UserModel{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		Address:      user.Address,
		DateOfBirth:  dob,
		RegisteredAt: user.RegisteredAt,
		Status:       user.Status,
	}

	if err := r.db.WithContext(context.Background()).Create(&model).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) FindByEmail(email string) (*entity.User, error) {
	if strings.TrimSpace(email) == "" {
		return nil, nil
	}

	var model usermodel.UserModel
	err := r.db.WithContext(context.Background()).
		Where("email = ?", strings.TrimSpace(email)).
		First(&model).Error

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, nil
	case err != nil:
		return nil, err
	}

	return modelToEntity(&model), nil
}

var _ repository.UserRepository = (*UserRepo)(nil)

func modelToEntity(model *usermodel.UserModel) *entity.User {
	if model == nil {
		return nil
	}

	var dob time.Time
	if model.DateOfBirth != nil {
		dob = *model.DateOfBirth
	}

	return &entity.User{
		ID:           model.ID,
		Name:         model.Name,
		Email:        model.Email,
		PhoneNumber:  model.PhoneNumber,
		Address:      model.Address,
		DateOfBirth:  dob,
		RegisteredAt: model.RegisteredAt,
		Status:       model.Status,
	}
}
````

## File: bookingapp/internal/infrastructure/repository/mysql/plan_repo_mysql.go
````go
package mysqlrepo

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	"bookingapp/internal/infrastructure/db/models"
	"context"
	"strings"

	"gorm.io/gorm"
)

type PlanRepo struct{ db *gorm.DB }

func NewPlanRepo(db *gorm.DB) repository.PlanRepository { return &PlanRepo{db: db} }

func (r *PlanRepo) FindByID(id int) (*entity.Plan, error) {
	var m models.PlanModel
	if err := r.db.WithContext(context.Background()).
		First(&m, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &entity.Plan{ID: m.ID, Name: m.Name, Keyword: m.Keyword, Price: m.Price}, nil
}

func (r *PlanRepo) SearchByKeyword(keyword string) ([]*entity.Plan, error) {
	ctx := context.Background()
	var list []models.PlanModel

	q := r.db.WithContext(ctx).Model(&models.PlanModel{})
	if strings.TrimSpace(keyword) != "" {
		kw := "%" + strings.TrimSpace(keyword) + "%"
		q = q.Where("LOWER(name) LIKE LOWER(?) OR LOWER(keyword) LIKE LOWER(?)", kw, kw)
	}
	if err := q.Order("id ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*entity.Plan, 0, len(list))
	for _, m := range list {
		copy := m
		out = append(out, &entity.Plan{ID: copy.ID, Name: copy.Name, Keyword: copy.Keyword, Price: copy.Price})
	}
	return out, nil
}
````

## File: bookingapp/internal/infrastructure/repository/mysql/reservation_repo_mysql.go
````go
package mysqlrepo

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	"bookingapp/internal/infrastructure/db/models"
	"context"

	"gorm.io/gorm"
)

type ReservationRepo struct{ db *gorm.DB }

func NewReservationRepo(db *gorm.DB) repository.ReservationRepository {
	return &ReservationRepo{db: db}
}

// DBのauto-incrementに委譲するのでNextIDは使わないが、interface満たすために実装
func (r *ReservationRepo) NextID() int { return 0 }

func (r *ReservationRepo) Save(res *entity.Reservation) (*entity.Reservation, error) {
	ctx := context.Background()
	m := models.ReservationModel{
		ID:       res.ID, // 0ならAUTO_INCREMENT
		UserID:   res.UserID,
		PlanID:   res.PlanID,
		Number:   res.Number,
		Checkin:  res.Checkin,
		Checkout: res.Checkout,
		Total:    res.Total,
	}
	if err := r.db.WithContext(ctx).Save(&m).Error; err != nil {
		return nil, err
	}
	// 生成されたIDを反映
	res.ID = m.ID
	return res, nil
}

func (r *ReservationRepo) FindByID(id int) (*entity.Reservation, error) {
	var m models.ReservationModel
	if err := r.db.WithContext(context.Background()).
		First(&m, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &entity.Reservation{
		ID:       m.ID,
		UserID:   m.UserID,
		PlanID:   m.PlanID,
		Number:   m.Number,
		Checkin:  m.Checkin,
		Checkout: m.Checkout,
		Total:    m.Total,
	}, nil
}

func (r *ReservationRepo) List() ([]*entity.Reservation, error) {
	var list []models.ReservationModel
	if err := r.db.WithContext(context.Background()).
		Order("id ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*entity.Reservation, 0, len(list))
	for _, m := range list {
		copy := m
		out = append(out, &entity.Reservation{
			ID:       copy.ID,
			UserID:   copy.UserID,
			PlanID:   copy.PlanID,
			Number:   copy.Number,
			Checkin:  copy.Checkin,
			Checkout: copy.Checkout,
			Total:    copy.Total,
		})
	}
	return out, nil
}

var _ repository.ReservationRepository = (*ReservationRepo)(nil)
````

## File: bookingapp/internal/interface/http/reservation_handler.go
````go
package httpi

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/usecase"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ReservationHandler struct {
	UC *usecase.ReservationUsecase
}

type createReq struct {
	UserID   string `json:"user_id"`
	PlanID   int    `json:"plan_id"`
	Number   int    `json:"number"`
	Checkin  string `json:"checkin"`  // "2025-10-12"
	Checkout string `json:"checkout"` // "2025-10-13"
}

type createResp struct {
	ID int `json:"id"`
}

type reservationView struct {
	ID       int    `json:"id"`
	UserID   string `json:"user_id"`
	PlanID   int    `json:"plan_id"`
	Number   int    `json:"number"`
	Checkin  string `json:"checkin"`
	Checkout string `json:"checkout"`
	Total    int    `json:"total"`
	Nights   int    `json:"nights"`
}

func (h *ReservationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var in createReq
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	ci, err1 := time.Parse("2006-01-02", in.Checkin)
	co, err2 := time.Parse("2006-01-02", in.Checkout)
	if err1 != nil || err2 != nil {
		http.Error(w, "invalid date format (yyyy-mm-dd)", http.StatusBadRequest)
		return
	}
	res, err := h.UC.Create(in.UserID, in.PlanID, in.Number, ci, co)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrInvalidUserID):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, usecase.ErrUserNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)
		case errors.Is(err, usecase.ErrInvalidDates):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, usecase.ErrInvalidNumber):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, usecase.ErrPlanNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}
	writeJSON(w, http.StatusOK, createResp{ID: res.ID})
}

func (h *ReservationHandler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/reservations/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	res, _ := h.UC.Get(id)
	if res == nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, http.StatusOK, toView(res))
}

func (h *ReservationHandler) List(w http.ResponseWriter, r *http.Request) {
	list, _ := h.UC.List()
	views := make([]reservationView, 0, len(list))
	for _, v := range list {
		views = append(views, toView(v))
	}
	writeJSON(w, http.StatusOK, views)
}

func (h *ReservationHandler) SearchPlans(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("keyword")
	plans, _ := h.UC.SearchPlans(q)
	type planView struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Keyword string `json:"keyword"`
		Price   int    `json:"price"`
	}
	out := make([]planView, 0, len(plans))
	for _, p := range plans {
		out = append(out, planView{ID: p.ID, Name: p.Name, Keyword: p.Keyword, Price: p.Price})
	}
	writeJSON(w, http.StatusOK, out)
}

// ここを *entity.Reservation にする（別型を作らない）
func toView(r *entity.Reservation) reservationView {
	return reservationView{
		ID:       r.ID,
		UserID:   r.UserID,
		PlanID:   r.PlanID,
		Number:   r.Number,
		Checkin:  r.Checkin.Format("2006-01-02"),
		Checkout: r.Checkout.Format("2006-01-02"),
		Total:    r.Total,
		Nights:   r.Nights(), // entity に既にあるメソッドを使う
	}
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
````

## File: bookingapp/internal/interface/http/user.go
````go
package httpi

import (
	"bookingapp/internal/usecase"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

type UserHandler struct {
	UC *usecase.UserUsecase
}

type registerUserReq struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
}

type registerUserResp struct {
	ID string `json:"id"`
}

type userView struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	Address      string `json:"address"`
	DateOfBirth  string `json:"date_of_birth"`
	RegisteredAt string `json:"registered_at"`
	Status       string `json:"status"`
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if h == nil || h.UC == nil {
		http.Error(w, "user lookup unavailable", http.StatusServiceUnavailable)
		return
	}

	id := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/users/"))
	if id == "" {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.UC.GetUser(id)
	if err != nil {
		if errors.Is(err, usecase.ErrUserInvalidInput) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.NotFound(w, r)
		return
	}

	view := userView{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		Address:      user.Address,
		DateOfBirth:  formatDate(user.DateOfBirth),
		RegisteredAt: user.RegisteredAt.Format(time.RFC3339),
		Status:       user.Status,
	}

	writeJSON(w, http.StatusOK, view)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if h == nil || h.UC == nil {
		http.Error(w, "user registration unavailable", http.StatusServiceUnavailable)
		return
	}

	var in registerUserReq
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	user, err := h.UC.Register(usecase.RegisterUserInput{
		Name:        in.Name,
		Email:       in.Email,
		PhoneNumber: in.PhoneNumber,
		Address:     in.Address,
		DateOfBirth: in.DateOfBirth,
	})
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrUserInvalidInput):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, usecase.ErrUserEmailAlreadyExists):
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	writeJSON(w, http.StatusCreated, registerUserResp{ID: user.ID})
}

func formatDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02")
}
````

## File: bookingapp/internal/usecase/user/user_uc.go
````go
package usecase

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
)

type UserUsecase struct {
	Users repository.UserRepository
}

// ユーザー情報取得
func (u *UserUsecase) GetUser(id string) (*entity.User, error) {
	return u.Users.Get(id)
}
````

## File: bookingapp/internal/usecase/reservation_uc.go
````go
package usecase

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidDates  = errors.New("invalid dates: checkout must be after checkin")
	ErrPlanNotFound  = errors.New("plan not found")
	ErrInvalidNumber = errors.New("number must be >= 1")
	ErrInvalidUserID = errors.New("invalid user id")
	ErrUserNotFound  = errors.New("user not found")
)

type ReservationUsecase struct {
	Users repository.UserRepository
	Plans repository.PlanRepository
	Resv  repository.ReservationRepository
}

// 　予約作成
func (u *ReservationUsecase) Create(userID string, planID, number int, checkin, checkout time.Time) (*entity.Reservation, error) {
	userID = strings.TrimSpace(userID)
	if userID == "" {
		return nil, ErrInvalidUserID
	}
	if _, err := uuid.Parse(userID); err != nil {
		return nil, ErrInvalidUserID
	}
	user, err := u.Users.Get(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	if !checkout.After(checkin) {
		return nil, ErrInvalidDates
	}
	if number < 1 {
		return nil, ErrInvalidNumber
	}
	plan, err := u.Plans.FindByID(planID)
	if err != nil {
		return nil, err
	}
	if plan == nil {
		return nil, ErrPlanNotFound
	}
	r := &entity.Reservation{
		ID:       u.Resv.NextID(),
		UserID:   user.ID,
		PlanID:   planID,
		Number:   number,
		Checkin:  checkin,
		Checkout: checkout,
	}
	//ドメイン層のメソッドを使って宿泊数を計算
	nights := r.Nights()
	//合計金額を計算してセット
	r.Total = plan.Price * number * nights
	//保存してID付きの予約情報を返す
	return u.Resv.Save(r)
}

// 予約取得
func (u *ReservationUsecase) Get(id int) (*entity.Reservation, error) {
	return u.Resv.FindByID(id)
}

// 予約一覧取得
func (u *ReservationUsecase) List() ([]*entity.Reservation, error) {
	return u.Resv.List()
}

// プラン検索
func (u *ReservationUsecase) SearchPlans(keyword string) ([]*entity.Plan, error) {
	return u.Plans.SearchByKeyword(keyword)
}
````

## File: bookingapp/internal/usecase/user_uc.go
````go
package usecase

import (
	"bookingapp/internal/domain/entity"
	"bookingapp/internal/domain/repository"
	"errors"
	"strings"
	"time"
)

var (
	ErrUserInvalidInput       = errors.New("invalid user input")
	ErrUserEmailAlreadyExists = errors.New("user email already exists")
)

type RegisterUserInput struct {
	Name        string
	Email       string
	PhoneNumber string
	Address     string
	DateOfBirth string
}

// ユースケース層からrepository層のinterfaceを使えるようにする
type UserUsecase struct {
	Users repository.UserRepository
	Now   func() time.Time
}

func (u *UserUsecase) Register(in RegisterUserInput) (*entity.User, error) {
	if u.Users == nil {
		return nil, errors.New("user repository is nil")
	}

	name := strings.TrimSpace(in.Name)
	email := strings.TrimSpace(in.Email)
	if name == "" || email == "" {
		return nil, ErrUserInvalidInput
	}

	existing, err := u.Users.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, ErrUserEmailAlreadyExists
	}

	var dob time.Time
	if v := strings.TrimSpace(in.DateOfBirth); v != "" {
		dob, err = time.Parse("2006-01-02", v)
		if err != nil {
			return nil, ErrUserInvalidInput
		}
	}

	now := time.Now
	if u.Now != nil {
		now = u.Now
	}

	user := &entity.User{
		Name:         name,
		Email:        email,
		PhoneNumber:  strings.TrimSpace(in.PhoneNumber),
		Address:      strings.TrimSpace(in.Address),
		DateOfBirth:  dob,
		RegisteredAt: now(),
		Status:       "active",
	}

	return u.Users.Create(user)
}

func (u *UserUsecase) GetUser(id string) (*entity.User, error) {
	if u.Users == nil {
		return nil, errors.New("user repository is nil")
	}

	trimmed := strings.TrimSpace(id)
	if trimmed == "" {
		return nil, ErrUserInvalidInput
	}

	return u.Users.Get(trimmed)
}
````

## File: bookingapp/docker-compose.yml
````yaml
version: "3.8"

services:
  mysql:
    image: mysql:8.0
    container_name: booking-mysql
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: booking
      MYSQL_USER: booking
      MYSQL_PASSWORD: bookingpass
      TZ: Asia/Tokyo
    ports:
      - "3306:3306"
    command:
      [
        "mysqld",
        "--default-authentication-plugin=mysql_native_password",
        "--character-set-server=utf8mb4",
        "--collation-server=utf8mb4_0900_ai_ci"
      ]
    volumes:
      - mysql_data:/var/lib/mysql
    networks:
      - booking-net

  api:
    build: .
    container_name: booking-api
    environment:
      DB_HOST: mysql
      DB_PORT: "3306"
      DB_USER: booking
      DB_PASS: bookingpass
      DB_NAME: booking
      PORT: "8080"
    depends_on:
      - mysql
    ports:
      - "8080:8080"
    networks:
      - booking-net

volumes:
  mysql_data: {}

networks:
  booking-net:
    driver: bridge
````

## File: bookingapp/go.mod
````
module bookingapp

go 1.24.0

require (
	gorm.io/driver/mysql v1.6.0
	gorm.io/gorm v1.31.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.3 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.30.0 // indirect
)
````

## File: bookingapp/mermaid.md
````markdown
erDiagram
    USERS {
      char36 id PK "users.id"
      varchar name
      varchar email UK "unique"
      varchar phone_number
      varchar address
      date date_of_birth
      datetime registered_at
      varchar status
      datetime created_at
      datetime updated_at
    }

    PLANS {
      int id PK
      varchar name
      varchar keyword
      int price
      datetime created_at
      datetime updated_at
    }

    RESERVATIONS {
      int id PK "reservations.id"
      char36 user_id FK "-> users.id"
      int plan_id FK "-> plans.id"
      int number
      date checkin
      date checkout
      int total
      datetime created_at
      datetime updated_at
    }

    USERS ||--o{ RESERVATIONS : "users.id = reservations.user_id"
    PLANS ||--o{ RESERVATIONS : "plans.id = reservations.plan_id"
````

## File: bookingapp/README.md
````markdown
# Booking App API

Go 言語と GORM を用いて宿泊予約を管理するシンプルな API サービスです。Clean Architecture を意識したレイヤ分割を採用し、ドメインロジックとインフラ依存コードを分離しています。

## 主な特徴
- MySQL（またはメモリリポジトリ）を利用したプラン・予約データ管理
- 予約作成時のバリデーション（宿泊日、人数、プラン存在チェック）
- RESTful なエンドポイント（プラン検索、予約 CRUD の一部）
- サーバ起動時の自動マイグレーションと初期データ投入

## ディレクトリ構成
```
.
├── cmd/api            # エントリーポイント（HTTP サーバ）
├── internal
│   ├── domain         # ドメインエンティティ & リポジトリインターフェース
│   ├── usecase        # ユースケース（アプリケーションサービス）
│   ├── interface/http # HTTP ハンドラ層
│   └── infrastructure # DB 接続・リポジトリ実装（MySQL / メモリ）
├── go.mod, go.sum
└── docker-compose.yml # MySQL 起動用定義
```

## アーキテクチャ概要
- `internal/domain/entity` にプラン (`Plan`) と予約 (`Reservation`) のドメインモデルを定義。`Reservation.Nights()` などのビジネスロジックをエンティティに寄せています。
- `internal/domain/repository` ではユースケースが依存するポート（インターフェース）を宣言。
- `internal/usecase/reservation_uc.go` はプラン検索や予約作成のアプリケーションロジックを担当し、入力バリデーションと料金計算を行います。
- `internal/interface/http` が HTTP リクエストを受け、ユースケースを呼び出して JSON を返却します。
- `internal/infrastructure` で具体的なアダプタを実装。`repository/mysql` は GORM を利用した永続化、`memory` はインメモリ実装です。

## 依存関係
- Go 1.24 以上
- MySQL 8 系（Docker Compose で起動可）
- ライブラリ
  - `gorm.io/gorm`
  - `gorm.io/driver/mysql`

## 環境変数
`cmd/api/main.go` では以下の環境変数を読み込み、未設定の場合はデフォルト値を使用します。

| 変数名   | デフォルト値  | 説明                 |
|----------|----------------|----------------------|
| `DB_HOST`| `127.0.0.1`    | MySQL ホスト         |
| `DB_PORT`| `3306`         | MySQL ポート         |
| `DB_USER`| `root`         | 接続ユーザー         |
| `DB_PASS`| `password`     | 接続パスワード       |
| `DB_NAME`| `booking`      | 使用するデータベース |

## 起動方法
1. 依存環境を用意
   ```bash
   docker compose up -d mysql
   ```
   > `docker-compose.yml` は root パスワード `password`、DB 名 `booking` で MySQL 8.0 を立ち上げます。

2. API サーバを起動
   ```bash
   go run ./cmd/api
   ```
   起動すると `:8080` で HTTP サーバが待ち受けます。

### マイグレーションとシード
サーバ起動時に以下が自動で実行されます。
- GORM の `AutoMigrate` による `plans` / `reservations` テーブル生成
- `plans` テーブルが空の場合、初期プラン 3 件を投入
  - 例: `ID=100, Name="富士プレミアム", Price=12000`

## ドメインロジック
- 予約作成 (`ReservationUsecase.Create`)
  - チェックイン < チェックアウト、人数 >= 1 を検証
  - 指定プランの存在確認（未存在時は `ErrPlanNotFound`）
  - 宿泊数（`Reservation.Nights()`）と人数、プラン単価から合計金額を算出
  - リポジトリ経由で保存し、生成された ID を返却
- 予約参照 (`Get`, `List`) とプラン検索 (`SearchPlans`) もユースケースを経由

## HTTP API
| メソッド | パス                | 説明                           |
|----------|---------------------|--------------------------------|
| `POST`   | `/reservations`     | 予約を新規作成                 |
| `GET`    | `/reservations`     | 予約一覧を取得                 |
| `GET`    | `/reservations/{id}`| 予約詳細を取得                 |
| `GET`    | `/plans`            | キーワードでプランを検索       |

### リクエスト/レスポンス例
**予約作成**
```bash
curl -X POST http://localhost:8080/reservations \
  -H 'Content-Type: application/json' \
  -d '{
        "plan_id": 100,
        "number": 2,
        "checkin": "2025-10-12",
        "checkout": "2025-10-14"
      }'
```
レスポンス
```json
{ "id": 1 }
```

**予約一覧**
```bash
curl http://localhost:8080/reservations
```
レスポンス（例）
```json
[
  {
    "id": 1,
    "plan_id": 100,
    "number": 2,
    "checkin": "2025-10-12",
    "checkout": "2025-10-14",
    "total": 48000,
    "nights": 2
  }
]
```

**プラン検索**
```bash
curl "http://localhost:8080/plans?keyword=富士"
```
レスポンス（例）
```json
[
  {
    "id": 100,
    "name": "富士プレミアム",
    "keyword": "富士 山 静岡",
    "price": 12000
  }
]
```

### エラーレスポンス
- リクエスト JSON / 日付フォーマット不正: `400 Bad Request`
- 宿泊日逆転・人数不足: `400 Bad Request`
- 存在しないプラン指定: `404 Not Found`
- その他予期しないエラー: `500 Internal Server Error`

## テストや拡張のヒント
- インメモリリポジトリ（`internal/infrastructure/memory`）を利用してユニットテストを書けます。
- バリデーション強化（例: 最大人数、予約重複チェック）や、キャンセル API 追加などの拡張が容易です。
- HTTP レイヤは `net/http` 標準ライブラリのままなので、Echo や Chi などに置き換える場合もユースケース層は流用可能です。

## ライセンス
本リポジトリにライセンス表記がない場合は、利用前に作成者へ確認してください。
````

## File: infra/bin/infra.ts
````typescript
#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib';
import { InfraStack } from '../lib/infra-stack';

const app = new cdk.App();
new InfraStack(app, 'InfraStack', {
  /* If you don't specify 'env', this stack will be environment-agnostic.
   * Account/Region-dependent features and context lookups will not work,
   * but a single synthesized template can be deployed anywhere. */

  /* Uncomment the next line to specialize this stack for the AWS Account
   * and Region that are implied by the current CLI configuration. */
  // env: { account: process.env.CDK_DEFAULT_ACCOUNT, region: process.env.CDK_DEFAULT_REGION },

  /* Uncomment the next line if you know exactly what Account and Region you
   * want to deploy the stack to. */
  // env: { account: '123456789012', region: 'us-east-1' },

  /* For more information, see https://docs.aws.amazon.com/cdk/latest/guide/environments.html */
});
````

## File: infra/lib/infra-stack.ts
````typescript
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
````

## File: infra/test/infra.test.ts
````typescript
// import * as cdk from 'aws-cdk-lib';
// import { Template } from 'aws-cdk-lib/assertions';
// import * as Infra from '../lib/infra-stack';

// example test. To run these tests, uncomment this file along with the
// example resource in lib/infra-stack.ts
test('SQS Queue Created', () => {
//   const app = new cdk.App();
//     // WHEN
//   const stack = new Infra.InfraStack(app, 'MyTestStack');
//     // THEN
//   const template = Template.fromStack(stack);

//   template.hasResourceProperties('AWS::SQS::Queue', {
//     VisibilityTimeout: 300
//   });
});
````

## File: infra/.gitignore
````
*.js
!jest.config.js
*.d.ts
node_modules

# CDK asset staging directory
.cdk.staging
cdk.out
````

## File: infra/.npmignore
````
*.ts
!*.d.ts

# CDK asset staging directory
.cdk.staging
cdk.out
````

## File: infra/cdk.json
````json
{
  "app": "npx ts-node --prefer-ts-exts bin/infra.ts",
  "watch": {
    "include": [
      "**"
    ],
    "exclude": [
      "README.md",
      "cdk*.json",
      "**/*.d.ts",
      "**/*.js",
      "tsconfig.json",
      "package*.json",
      "yarn.lock",
      "node_modules",
      "test"
    ]
  },
  "context": {
    "@aws-cdk/aws-signer:signingProfileNamePassedToCfn": true,
    "@aws-cdk/aws-ecs-patterns:secGroupsDisablesImplicitOpenListener": true,
    "@aws-cdk/aws-lambda:recognizeLayerVersion": true,
    "@aws-cdk/core:checkSecretUsage": true,
    "@aws-cdk/core:target-partitions": [
      "aws",
      "aws-cn"
    ],
    "@aws-cdk-containers/ecs-service-extensions:enableDefaultLogDriver": true,
    "@aws-cdk/aws-ec2:uniqueImdsv2TemplateName": true,
    "@aws-cdk/aws-ecs:arnFormatIncludesClusterName": true,
    "@aws-cdk/aws-iam:minimizePolicies": true,
    "@aws-cdk/core:validateSnapshotRemovalPolicy": true,
    "@aws-cdk/aws-codepipeline:crossAccountKeyAliasStackSafeResourceName": true,
    "@aws-cdk/aws-s3:createDefaultLoggingPolicy": true,
    "@aws-cdk/aws-sns-subscriptions:restrictSqsDescryption": true,
    "@aws-cdk/aws-apigateway:disableCloudWatchRole": true,
    "@aws-cdk/core:enablePartitionLiterals": true,
    "@aws-cdk/aws-events:eventsTargetQueueSameAccount": true,
    "@aws-cdk/aws-ecs:disableExplicitDeploymentControllerForCircuitBreaker": true,
    "@aws-cdk/aws-iam:importedRoleStackSafeDefaultPolicyName": true,
    "@aws-cdk/aws-s3:serverAccessLogsUseBucketPolicy": true,
    "@aws-cdk/aws-route53-patters:useCertificate": true,
    "@aws-cdk/customresources:installLatestAwsSdkDefault": false,
    "@aws-cdk/aws-rds:databaseProxyUniqueResourceName": true,
    "@aws-cdk/aws-codedeploy:removeAlarmsFromDeploymentGroup": true,
    "@aws-cdk/aws-apigateway:authorizerChangeDeploymentLogicalId": true,
    "@aws-cdk/aws-ec2:launchTemplateDefaultUserData": true,
    "@aws-cdk/aws-secretsmanager:useAttachedSecretResourcePolicyForSecretTargetAttachments": true,
    "@aws-cdk/aws-redshift:columnId": true,
    "@aws-cdk/aws-stepfunctions-tasks:enableEmrServicePolicyV2": true,
    "@aws-cdk/aws-ec2:restrictDefaultSecurityGroup": true,
    "@aws-cdk/aws-apigateway:requestValidatorUniqueId": true,
    "@aws-cdk/aws-kms:aliasNameRef": true,
    "@aws-cdk/aws-kms:applyImportedAliasPermissionsToPrincipal": true,
    "@aws-cdk/aws-autoscaling:generateLaunchTemplateInsteadOfLaunchConfig": true,
    "@aws-cdk/core:includePrefixInUniqueNameGeneration": true,
    "@aws-cdk/aws-efs:denyAnonymousAccess": true,
    "@aws-cdk/aws-opensearchservice:enableOpensearchMultiAzWithStandby": true,
    "@aws-cdk/aws-lambda-nodejs:useLatestRuntimeVersion": true,
    "@aws-cdk/aws-efs:mountTargetOrderInsensitiveLogicalId": true,
    "@aws-cdk/aws-rds:auroraClusterChangeScopeOfInstanceParameterGroupWithEachParameters": true,
    "@aws-cdk/aws-appsync:useArnForSourceApiAssociationIdentifier": true,
    "@aws-cdk/aws-rds:preventRenderingDeprecatedCredentials": true,
    "@aws-cdk/aws-codepipeline-actions:useNewDefaultBranchForCodeCommitSource": true,
    "@aws-cdk/aws-cloudwatch-actions:changeLambdaPermissionLogicalIdForLambdaAction": true,
    "@aws-cdk/aws-codepipeline:crossAccountKeysDefaultValueToFalse": true,
    "@aws-cdk/aws-codepipeline:defaultPipelineTypeToV2": true,
    "@aws-cdk/aws-kms:reduceCrossAccountRegionPolicyScope": true,
    "@aws-cdk/aws-eks:nodegroupNameAttribute": true,
    "@aws-cdk/aws-ec2:ebsDefaultGp3Volume": true,
    "@aws-cdk/aws-ecs:removeDefaultDeploymentAlarm": true,
    "@aws-cdk/custom-resources:logApiResponseDataPropertyTrueDefault": false,
    "@aws-cdk/aws-s3:keepNotificationInImportedBucket": false,
    "@aws-cdk/core:explicitStackTags": true,
    "@aws-cdk/aws-ecs:enableImdsBlockingDeprecatedFeature": false,
    "@aws-cdk/aws-ecs:disableEcsImdsBlocking": true,
    "@aws-cdk/aws-ecs:reduceEc2FargateCloudWatchPermissions": true,
    "@aws-cdk/aws-dynamodb:resourcePolicyPerReplica": true,
    "@aws-cdk/aws-ec2:ec2SumTImeoutEnabled": true,
    "@aws-cdk/aws-appsync:appSyncGraphQLAPIScopeLambdaPermission": true,
    "@aws-cdk/aws-rds:setCorrectValueForDatabaseInstanceReadReplicaInstanceResourceId": true,
    "@aws-cdk/core:cfnIncludeRejectComplexResourceUpdateCreatePolicyIntrinsics": true,
    "@aws-cdk/aws-lambda-nodejs:sdkV3ExcludeSmithyPackages": true,
    "@aws-cdk/aws-stepfunctions-tasks:fixRunEcsTaskPolicy": true,
    "@aws-cdk/aws-ec2:bastionHostUseAmazonLinux2023ByDefault": true,
    "@aws-cdk/aws-route53-targets:userPoolDomainNameMethodWithoutCustomResource": true,
    "@aws-cdk/aws-elasticloadbalancingV2:albDualstackWithoutPublicIpv4SecurityGroupRulesDefault": true,
    "@aws-cdk/aws-iam:oidcRejectUnauthorizedConnections": true,
    "@aws-cdk/core:enableAdditionalMetadataCollection": true,
    "@aws-cdk/aws-lambda:createNewPoliciesWithAddToRolePolicy": false,
    "@aws-cdk/aws-s3:setUniqueReplicationRoleName": true,
    "@aws-cdk/aws-events:requireEventBusPolicySid": true,
    "@aws-cdk/core:aspectPrioritiesMutating": true,
    "@aws-cdk/aws-dynamodb:retainTableReplica": true,
    "@aws-cdk/aws-stepfunctions:useDistributedMapResultWriterV2": true,
    "@aws-cdk/s3-notifications:addS3TrustKeyPolicyForSnsSubscriptions": true,
    "@aws-cdk/aws-ec2:requirePrivateSubnetsForEgressOnlyInternetGateway": true,
    "@aws-cdk/aws-s3:publicAccessBlockedByDefault": true,
    "@aws-cdk/aws-lambda:useCdkManagedLogGroup": true
  }
}
````

## File: infra/jest.config.js
````javascript
module.exports = {
  testEnvironment: 'node',
  roots: ['<rootDir>/test'],
  testMatch: ['**/*.test.ts'],
  transform: {
    '^.+\\.tsx?$': 'ts-jest'
  }
};
````

## File: infra/package.json
````json
{
  "name": "infra",
  "version": "0.1.0",
  "bin": {
    "infra": "bin/infra.js"
  },
  "scripts": {
    "build": "tsc",
    "watch": "tsc -w",
    "test": "jest",
    "cdk": "cdk"
  },
  "devDependencies": {
    "@types/jest": "^29.5.14",
    "@types/node": "22.7.9",
    "jest": "^29.7.0",
    "ts-jest": "^29.2.5",
    "aws-cdk": "2.1030.0",
    "ts-node": "^10.9.2",
    "typescript": "~5.6.3"
  },
  "dependencies": {
    "aws-cdk-lib": "2.215.0",
    "constructs": "^10.0.0"
  }
}
````

## File: infra/README.md
````markdown
# Welcome to your CDK TypeScript project

This is a blank project for CDK development with TypeScript.

The `cdk.json` file tells the CDK Toolkit how to execute your app.

## Useful commands

* `npm run build`   compile typescript to js
* `npm run watch`   watch for changes and compile
* `npm run test`    perform the jest unit tests
* `npx cdk deploy`  deploy this stack to your default AWS account/region
* `npx cdk diff`    compare deployed stack with current state
* `npx cdk synth`   emits the synthesized CloudFormation template
````

## File: infra/tsconfig.json
````json
{
  "compilerOptions": {
    "target": "ES2022",
    "module": "NodeNext",
    "moduleResolution": "NodeNext",
    "lib": [
      "es2022"
    ],
    "declaration": true,
    "strict": true,
    "noImplicitAny": true,
    "strictNullChecks": true,
    "noImplicitThis": true,
    "alwaysStrict": true,
    "noUnusedLocals": false,
    "noUnusedParameters": false,
    "noImplicitReturns": true,
    "noFallthroughCasesInSwitch": false,
    "inlineSourceMap": true,
    "inlineSources": true,
    "experimentalDecorators": true,
    "strictPropertyInitialization": false,
    "skipLibCheck": true,
    "typeRoots": [
      "./node_modules/@types"
    ]
  },
  "exclude": [
    "node_modules",
    "cdk.out"
  ]
}
````

## File: fix.md
````markdown
ユーザー会員登録

 - bookingapp/internal/domain/repository/repository.go:17 に UserRepository を追加し、ユースケースが依存するインターフェースを定義。
  - bookingapp/internal/usecase/user_uc.go:29 で登録ロジックを実装し、入力バリデーション・日付パース・重複メール検出を行って UserRepository へ委譲。
  - bookingapp/internal/infrastructure/repository/mysql/user/user_repo_mysql.go:24 に MySQL 実装を追加し、UUID 発行・DateOfBirth の NULL 対応・既存メール検索を実装。
  - bookingapp/internal/interface/http/user.go:26 にユーザーハンドラを作成し、HTTP からユースケースを呼び出して適切な HTTP ステータスを返却。
  - bookingapp/cmd/api/main.go:42 で新しいハンドラ/ユースケース/リポジトリを初期化し、POST /register を UserHandler にバインド。
````
