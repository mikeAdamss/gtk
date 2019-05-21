
# gtk

gtk (good to know) is a simple http request-response controller intended to allow the convenient control of large numbers (or rarely used) rest based severless functions (though it'll work with any restful GET endpoints).

It's a go wrapper around an user controlled inventory of endpoints, their descriptions and their identifers. The parts that're "good to know" is what functions are availible to you, whether a called function succeeded and what was returned in the response.


## How it works

The endpoints are defined via a user controlled inventory. Each item in the inventory consists of:

- a keyword to call the endpoint.
- a description of what it does and what it returns.
- its url.

and ..optionally
- any headers that need to be added to the request, either hard coded or from an environmental variable.
- the chosen parser (if any), to handle the response body returned.


## Installation and Setup

- Clone this repo.
- Set the env variable `GTK_INVENTORY` pointing to the included `inventory.yml`.
- Navigate to the cloned repo and `go install`.


## Example commands

default:
- `gtk`: lists all endpoints gtk knows about with the description of what they do.
- `gtk -call=<NAME-OF_ENDPOINT>`: call the endpoint in question.

NOTE - to use the "ports" example you'll need update the url field in the yaml file (I cant put it on github).


## Defining a new inventory item

To add an endpoint to to the inventory add an entry as per the below to your `inventory.yaml` file.

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

The response_parses will be build up over time as needed. The "arrayOfLines" one shown expects to be returned a response body that can be unmarashalled into an array of strings, these strings are then printed to the terminal in order.

If no `response_parser` is specified gdk will print the response status code.
