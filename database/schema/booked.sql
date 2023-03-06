CREATE TABLE Booked (
    ssn INTEGER,
    train_no INTEGER,
    category VARCHAR(10),
    status VARCHAR(10),
    FOREIGN KEY (ssn) REFERENCES Passenger (ssn),
    FOREIGN KEY (train_no) REFERENCES Train (train_no)
);
