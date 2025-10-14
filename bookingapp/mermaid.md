erDiagram
    USERS {
      int id PK
      varchar name
      varchar email UK
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
      int id PK
      int user_id FK
      int plan_id FK
      date checkin
      date checkout
      int number
      int total
      datetime created_at
      datetime updated_at
    }

    USERS ||--o{ RESERVATIONS : has
    PLANS ||--o{ RESERVATIONS : includes
