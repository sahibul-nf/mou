### Hasil analisis entity:

- Users
- Campaigns
- Campaign Images
- Transactions
\
\
\\

```plantuml
@startuml 
    entity Users {
        * id : INT
        --
        name : VARCHAR
        email : VARCHAR
        password_hash : VARCHAR
        occupation : VARCHAR
        avatar_file_name : VARCHAR
        role : VARCHAR
        token : VARCHAR
        created_at : DATETIME
        updated_at : DATETIME
    }

    entity Campaigns {

    }

    entity CampaignImages {

    }

    entity Transactions {

    }

Users ||--o{ Campaigns : Create
Campaigns ||--|{ CampaignImages : Has
Users ||--o{ Transactions : Has
Transaction }o--|| Campaigns : Has

@enduml
```
