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
