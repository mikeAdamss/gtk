
# gdk (good to know)

This is a go wrapper intended for controlling serverless functions. The bit that's "good to know" is whether the function succeeded
and what was returned - this response gets cleaned up and printed to your terminal, formatted via the specified parser.

...wut?
Try it, it's fab.


## How it works

The functions are defined via a user controlled inventory. Each item in the inventory consists of:

- a keyword to call the function.
- a description of what it does and what it returns.
- its url.

and ..optionally
- any headers that need to be added to the request, either hard coded or from an environmental variable.


## Installation and Setup

- Clone this repo.
- Set the env variable `GTK_INVENTORY` pointing to the included `inventory.yml`.
- Navigate to the cloned repo and `go install`.


## Example commands

default:
- `gtk`: lists all serverless functions gtk knows about with the description of what they do.

inventory defined:
- `gtk -call=ports`: a simple example function I've included. Scrapes then prints all services and ports listed in the dp repo.


# Defining an inventory item (i.e adding more serverless functions)

Create the serverless function, add an entry as per the below example to your inventory.yaml file.

```yaml
functions:

    # An example using two request headers, one hard coded, one taken from an environmental variables.
    - name: "example"
      description: "This is just the example, a description goes here."
      url: "www.some-variation-of-serveless-function-url.com/probably-google-or-aws/example"
      headers:
         - key: "foo-key"
           value: "foo"
           env: false
         - name: "baa_key"
           value: "BAA_KEY"
           env: true
      response_parser: "arrayOfLines"
```

The response_parses will be build up over time. The "arrayOfLines" one shows, expects to be returned a response body that can be
unmarashalled into an array of strings, which are then printed to your terminal.

If no response_parser is specified gdk will print the response status code.

** OTHER - All methods are GET and I haven't done the logic for adding headers yet!**