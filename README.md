Go CLI App
==========

Go CLI app to compare the share values of AAPL and MSFT (10)shares and out put the profit or the loss made

Command to run the app
```./cliapp checkShareValue```

# Info
App uses the Cobra Plugin to implement a CLI

# Assumptions:
Number of shares = 10
Shares of only AAPL and MSFT
These values live in the config for the test

# TBD
1. Config file path is hardcoded, best to give it in as a flag, with a default in case it is 2. not given
3. Add validation to config file
4. app/calculations/pnl.go is outputting string, ideally just calculations and another "reporting" interface to convert to useful messages


