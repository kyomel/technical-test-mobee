-- Create the Users table
CREATE TABLE IF NOT EXISTS Users (
    UserID SERIAL PRIMARY KEY,
    Username VARCHAR(50) UNIQUE NOT NULL,
    Password VARCHAR(100) NOT NULL,
    Email VARCHAR(100) UNIQUE NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the Wallets table
CREATE TABLE IF NOT EXISTS Wallets (
    WalletID SERIAL PRIMARY KEY,
    UserID INT,
    Balance DECIMAL(15, 2) DEFAULT 0.00,
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

-- Create the Transactions table
CREATE TABLE IF NOT EXISTS Transactions (
    TransactionID SERIAL PRIMARY KEY,
    TransactionType VARCHAR(20) NOT NULL CHECK (TransactionType IN ('Deposit', 'Withdraw', 'Transfer')),
    Amount DECIMAL(15, 2) NOT NULL,
    SenderID INT,
    ReceiverID INT,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (SenderID) REFERENCES Users(UserID),
    FOREIGN KEY (ReceiverID) REFERENCES Users(UserID)
);

-- Create the Ledger table
CREATE TABLE IF NOT EXISTS Ledger (
    LedgerID SERIAL PRIMARY KEY,
    TransactionID INT,
    SenderID INT,
    ReceiverID INT,
    Amount DECIMAL(15, 2) NOT NULL,
    BalanceBefore DECIMAL(15, 2) NOT NULL,
    BalanceAfter DECIMAL(15, 2) NOT NULL,
    Timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (TransactionID) REFERENCES Transactions(TransactionID),
    FOREIGN KEY (SenderID) REFERENCES Users(UserID),
    FOREIGN KEY (ReceiverID) REFERENCES Users(UserID)
);