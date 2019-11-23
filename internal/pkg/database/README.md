# How to use the `database` package

You can get full control over the database connection by calling `database.GetConnection()`, this function gives you the sqlx object used to call SQL queries.

The database has three formatting rules for the table and column names:

1. Table names have UPPERCASE names.
2. Column names have lowercase names.
3. Both table names and column names use underscores to split words.

## Example of usage

Look at `user_test.go` for example of usage. This file also shows how to ROLLBACK changes made to the database. ROLLBACK works in cohesion with START TRANSACTION.

## How to connect the database to GoLand

1. On the far right of the screen there is a button called "Database". Click it!

2. Navigate to adding a database connection, select "MariaDB".

3. Insert all the database information found in .cache/db-config.json - including the password.

4. Download the drivers.

5. Test the connection.

6. Does it work? Great, continue the guide - if it doesn't work; ask for help.

7. In the previously opened "Database" tab - sync the database.

8. ???

9. Profit

## `database` package

### USER

This table contains GOOGLE authentication data. Nothing related to lichess.

### RESULT

This table contains match data, and results.

Outcome can be 3 values - "win", "draw" or "loss".

### RESULT_PLATFORM_ELO

This table connect the result to an ELO platform, for example lichess.org.

### PLATFORM

This table contains data about platforms we utilize, like Google and lichess.org.

### PLATFORM_AUTH

Which platforms are authentication platforms?

### PLATFORM_ELO

Which platforms are ELO platforms?

### PLATFORM_ENDPOINT

Insert the endpoints for each platform here - use custom names to remember what is what.

### PLATFORM_USER

Connect users with a platform.

AccessToken is used to gain access and identify to the PLATFORM.

VerificationKey is used to identify the USER with the PLATFORM identity.

### LEAGUE

When importing a lichess.org team, create them as a LEAGUE.

So `Storbukk Sjakklub` is a LEAGUE.

### LEAGUE_SEASON

Each league can have many seasons where ELO is reset.

For our purposes just create a season with the name `CONTINUOUS` when creating the LEAGUE.

### GROUP

Groups are the players playing in the SEASON in the LEAGUE. For chess there will only be one player in each group, so when importing from lichess.org just create a GROUP for every single lichess.org player within the imported teams. For example a group can be the player `hyge`.

### GROUP_USER

Connects a USER with a GROUP. This connection is made possible when using the PLATFORM_USER VerificationKey, connecting the identity of the lichess.org user with our USER - thus making it possible to connect a USER with a GROUP.
