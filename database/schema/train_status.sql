CREATE TABLE Train_Status (
    train_no integer,
    date date,
    ps_available integer,
    gs_available integer,
    ps_occupied integer,
    gs_occupied integer, 
    FOREIGN KEY (train_no) REFERENCES Train (train_no)
);
