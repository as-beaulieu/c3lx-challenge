# c3lx-challenge

## disclaimers

POST call to accept challenge returns accepted state before the call. 
Another call to GET challenges will reflect the change that was made.

## go mod support

`go mod init`

## SQL Scripts used

```sql
INSERT INTO ablue.challenges
(name, description, fullname)
VALUES
(
 'steps',
 'Get at least 10,000 steps a day for a month',
 'Get to steppin'
);

SELECT * FROM ablue.challenges;

INSERT INTO ablue.challenges
(name, description, fullname)
VALUES
(
    'eat',
    'Eat two healthy meals a day for two weeks',
    'You are what you eat'
);

INSERT INTO ablue.challenges
(name, description, fullname)
VALUES
(
    'sleep',
    'Average seven hours of sleep for five nights in a row',
    'Rip Van Winkle'
);

alter table ablue.challenges
	add accepted bool;
```