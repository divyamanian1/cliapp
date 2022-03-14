Go CLI App
==========

Go CLI app to compare the share values of AAPL and MSFT (10)shares and out put the profit or the loss made

Command to run the app
```./cliapp checkShareValue MSFT AAPL```

# Info
App uses the Cobra Plugin to implement a CLI

# Assumptions:
Number of shares = 10
Shares of only AAPL and MSFT
These values live in the config for the test

# TBD
1. Config file with API key instead of using a constant
2. Validation on the symbols and their count
3. Add validation to config file
4. app/calculations/pnl.go is outputting string, ideally just calculations and another "reporting" interface to convert to useful messages


