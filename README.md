# c3lx-challenge

Written with golang 1.15

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
```