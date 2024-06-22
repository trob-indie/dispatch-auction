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

### PUT /api/auction
Starts an auction with the given ID. If an auction has not yet been started, this will start the auction with all registered users. If the auction has been completed, the historical auction data will be fetched from the database.

**Request**:
```
{
    "auction_id": uuid,
    "owner_id": uuid
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

### POST /api/auction/register
Registers a bidder to an auction. IF the user is already registered as a bidder for this auction, the existing auction-registration record is returned
**Request Body:**
```
{
    "auction_id": uuid,
    "bidder_id": uuid
}
```

**Response Body:**
```
{
    "registration_id": uuid,
    "auction_id": uuid,
    "bidder_id": uuid,
    "created_at": timestamp,
    "updated_at": timestamp
}
```

## Database Schema
The REST API provides an interface for managing resources in a PostgreSQL database.

### table: users
Stores a record for each user.
**Columns:**
id (PK): uuid
username: string
password: hash(string)
created_at: timestamp
updated_at: timestamp

### table: auctions
Stores a record for each auction.
**Columns:**
id (PK): uuid
owner_id: uuid
title: string
description: string
images: json_blob
winning_user_id: uuid
winning_bid: int32
created_at: timestamp
updated_at: timestamp

### table: auction_registrations
Stores a record for each bidder in an auction.
**Columns:**
id (PK): uuid
auction_id (FK): uuid
bidder_id (FK): uuid
created_at: timestamp
updated_at: timestamp

### table: bids
Stores a record each time a user bids in an auction
**Columns:**
id (PK): uuid
auction_id (FK): uuid
bidder_id (FK): uuid
max_bid: int32
current_bid: int32
auto_increment: int32
created_at: timestamp


## Future Enhancements
* Hash the password before storing it in the database
* Use UUIDs from the github.com/google/uuid package
* Authentication and Authorization should be considered. A user could be issued a JWT which contains authorization scopes relevant to each endpoint.
    - In this design, a user can be an owner or a bidder, but maybe we would want to restrict access based on user type
    - Ideally, only the owner of the auction should be able to start it
* Add a list action endpoint for the front-end to display to the bidders
* Return proper HTTP errors
* Split the auction_registration table into 2 tables -- 1 for just registration, and the other for bids
* Give a bidder the ability to un-enroll from an auction
* Give the auction owner the ability to cancel the auction
* Give the auction owner the ability to set a start time for the auction instead of manually starting.
* Use an actual postgres database that can be persisted between runs