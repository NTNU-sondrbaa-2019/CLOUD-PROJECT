# gr8elo.com

An application for measuring ELO within a small collections of players or groups of players.

## Group `8`: `GR8`

* `Sondre Benjamin Aasen` (`sondrbaa`)
* `Adrian Emil Chambe-Eng` (`adrianec`)
* `Marcus Hyge`  (`marcushy`)
* `Vebjørn Fonstad Leiros` (`vebjorfl`)
* `Sindre Østrem` (`sindrost`)

**Project name:** `gr8elo`

#### Project idea

Users login using `Google Identity Platform` and we collect data from `lichess.org API`. Local session tokens are stored in our cache, all of the remaining data is stored in our `AWS RDS`. The main function of the solution is to collect results of matches within lichess.org Teams and calculating local ELO within the teams. After we've implemented this functionality we'll look into expanding the idea by analysing other data from the `lichess.org API` - such as analysing ELO growth by played matches, winrates of nation vs nation and other unexplored statistics and analysing tools.

We plan on using this project for student activity groups. Such as `NTNUI-Gjøvik Sjakk` and `NTNUI-Gjøvik Esport`. Therefor we're trying to design everything to be as flexible as possible now - so that we can more easily continue the project after the Cloud Technologies project is over.

#### Technologies

* `GO`
* `Heroku`
* `Google Identity Platform`
* `lichess.org API`
* `AWS RDS`
* `Cronjobs`

#### Resources

* https://aws.amazon.com
* https://lichess.org/api
* https://developers.google.com/identity
* https://heroku.com
* https://namecheap.com

## Report

### What changed from our original project idea?

TODO What has changed?

TODO Marcus

TODO Sindre

TODO Adrian

TODO Vebjørn

### What did we not achieve with the project?

TODO What was not achieved?

We have some bugs and some other features that we did not manage to finish in time. We also could need some more tests that we did not achieve to do within the time frame we had. 

Wanted modularity, the rating package could be split into three packages.
One for the lichess.org api calls, one for calculations, and the last for combining theese two and do the database queries.

The rating package has a bug at the moment when it fetches the current elo of a player in the database, elo is always set to 1500 before a new calculation is done to the players elo.
This makes the use case of this package unusable as the elo always will be 1500 +- rating change of one match.

TODO Adrian

TODO Vebjørn

### Reflections

We needed more time to design the project together instead of going straight to development. The scope of the project was too large for our group; since we're all very busy students.

A solution that would help with the development of the project for our group would be this: in the early phase of the development; create packages with all structures, functions and the general layout of our project - without implementing anything - as empty shells. Just empty interfaces. This would help everybody understand each module with greater ease.

### New Experiences and Knowledge

TODO What did we learn?

Learned how to deal with NDJSON structs, more about writing tests and how to work with another person on the same tasks. Time management is also something we can take as an big experience from this project. Its easy to dream big, but things takes time, and most of the cases it takes more time that you think and plan.  

TODO Sindre

TODO Adrian

We worked with an `AWS RDS` database and learned how to setup the instance, connect to it and run SQL from Go. This will be useful in the future since persitent data is useful for every project.

TODO Vebjørn

### Work Hours

The group accumulated a total of 200 hours of work hours.
