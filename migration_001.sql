CREATE DATABASE dispatch;

CREATE TABLE IF NOT EXISTS users (
    id: VARCHAR(16),
    username: VARCHAR(32),
    password: VARCHAR(32),
    created_at: TIMESTAMP,
    updated_at: TIMESTAMP,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS auctions (
    id: VARCHAR(16),
    owner_id: VARCHAR(16),
    title: VARCHAR(64),
    description: TEXT,
    images: JSONB,
    winning_user_id: VARCHAR(16),
    winning_bid: INT
    created_at: TIMESTAMP,
    updated_at: TIMESTAMP,
    PRIMARY KEY(id),
    CONSTRAINT fk_owner
      FOREIGN KEY(owner_id) 
        REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS auction_registrations (
    id: VARCHAR(16),
    auction_id: VARCHAR(16),
    bidder_id: VARCHAR(16),
    max_bid: INT,
    initial_bid: INT,
    auto_increment: INT,
    PRIMARY KEY(id),
    created_at: TIMESTAMP,
    updated_at: TIMESTAMP,
    CONSTRAINT fk_auction
    FOREIGN KEY(auction_id) 
        REFERENCES auctions(id),
    CONSTRAINT fk_bidder
      FOREIGN KEY(bidder_id) 
        REFERENCES users(id)
);
