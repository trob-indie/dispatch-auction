# Dispatch Auction Take-Home
This repo is my implementation of the Dispatch take-home challenge.

## REST API
This Go module exposes multiple REST endpoints for managing users and their relationships with auctions.

### POST /api/user
Creates a new user.

**Request Body**: 
```
{
    "username": string,
    "password": string
}
```

**Response Body**:
```
{
    "id": uuid,
    "username": string,
    "created_at": timestamp,
    "updated_at": timestamp
}
```

### POST /api/auction
Creates a new auction. Users can register to an auction after it has been created

**Request Body**: 
```
{
    "owner_id": uuid,
    "title": string,
    "description": string,
    "images": [
        {
            "url": string,
            "size": {
                width: int32,
                height: int32
            },
            "resolution": string
        },
        ...
    ]
}
```

**Response Body**:
```
{
    "id": uuid,
    "owner_id": uuid,
    "title": string,
    "description": string,
    "images": [
        {
            "url": string,
            "size": {
                width: int32,
                height: int32
            },
            "resolution": string
        },
        ...
    ],
    "created_at": timestamp,
    "updated_at": timestamp
}
```

### GET /api/auction/:auction_id:
Gets an auction by ID. If an auction has not yet been started, this will start the auction with all registered users. If the auction has been completed, the historical auction data will be fetched from the database.

**Response Body**:
```
{
    "id": uuid,
    "owner_id": uuid,
    "title": string,
    "description": string,
    "images": [
        {
            "url": string,
            "size": {
                width: int32,
                height: int32
            },
            "resolution": string
        },
        ...
    ],
    "bidders": [
        uuid,
        uuid
    ],
    "winning_user_id": nullable uuid,
    "winning_bid": int32
    "created_at": timestamp,
    "updated_at": timestamp
}
```

## Database Schema
The REST API provides an interface for managing resources in a PostgreSQL database.

### table: auctions
Stores a record for each auction.
**Columns:**
id (PK): uuid
name: string
description: string
images: json_blob
owner_id: uuid
winning_user_id: uuid
winning_bid: int32
created_at: timestamp
updated_at: timestamp

### table: users
Stores a record for each user.
**Columns:**
id (PK): uuid
username: string
password: hash(string)
created_at: timestamp
updated_at: timestamp

### table: auction_registrations
Stores a record for each auction-user relationship.
**Columns:**
id (PK): uuid
user_id (FK): uuid
auction_id (FK): uuid
max_bid: int32
current_bid: int32
auto_increment: int32
created_at: timestamp
updated_at: timestamp


## Future Enhancements
* Hash the password before storing it in the database
* Use real UUIDs using the github.com/google/uuid package
* Authentication and Authorization should be considered. A user could be issued a JWT which contains authorization scopes relevant to each endpoint.
    - In this design, a user can be an actioneer or a bidder, but maybe we would want to restrict access based on user type
* Add a list action endpoint for the front-end to display to the bidders