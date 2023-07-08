CREATE TABLE personalities(
    id serial primary key,
    name varchar,
    history varchar
);

INSERT INTO personalities(name, history) VALUES
('Linus Torvalds', 'Linus Benedict Torvalds is a Finnish software engineer who is the creator and, historically, the lead developer of the Linux kernel, used by Linux distributions and other operating systems such as Android. He also created the distributed version control system Git.'),
('Jon Hall', 'Jon "maddog" Hall is the board chair for the Linux Professional Institute');