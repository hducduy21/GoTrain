module example/hello

go 1.23.4

require example.com/greetings v1.1.0
require example.com/db v1.1.0

replace example.com/greetings => ../greetings

replace example.com/db => ../db
