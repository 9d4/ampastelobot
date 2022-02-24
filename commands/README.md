### /commands

`package commands`

Every command has it's own file, but not it's own package. This package is where we can examine the command that received from matcher `/matchers`. Then we can do more logic with api from `/app`, there we store anything complex in single package. 

Keep in mind, this package is not for complex logic. Make complex thing in `/app`.
