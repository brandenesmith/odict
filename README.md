# odict
A command line Oxford Dictionary utility.

# Getting Started 
Using the odict command line utility requires an API id and an API key.
You can obtain these items by visiting the [Oxford Dictionary Developer Portal](https://developer.oxforddictionaries.com).
Once you have obtained the id and key, you should set the environment variables 
`ODICTAPIID` and `ODICTAPIKEY` with the id and key respectively. To set the environment variables,
you can simply add the following lines to your `.bash_profile`:
```bash
export ODICTAPIID=<your id>
export ODICTAPIKEY=<your key>
```
After these lines are added, restart your terminal for the exports to become effective. 