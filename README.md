
# gtk

gtk (good to know) is a simple serverless function controller.

It's a go wrapper around an user controlled inventory of aws lambda and google cloud functions. The bit that's "good to know" is whether the function succeeded and what was returned.


## How it works

The functions are defined via a user controlled inventory. Each item in the inventory consists of:

- a keyword to call the function.
- a description of what it does and what it returns.
- its url.

and ..optionally
- any headers that need to be added to the request, either hard coded or from an environmental variable (note - not implemented yet)
- the chosen parser (if any), to handle the response body returned from the function.


## Installation and Setup

- Clone this repo.
- Set the env variable `GTK_INVENTORY` pointing to the included `inventory.yml`.
- Navigate to the cloned repo and `go install`.


## Example commands

default:
- `gtk`: lists all serverless functions gtk knows about with the description of what they do.

inventory defined:
- `gtk -call=ports`: a simple example, scrapes then prints all services and ports listed in the dp repo.

NOTE - to use the "ports" example you'll need update the url (I cant put that on github).


# Defining an inventory item (i.e adding more serverless functions)

To add a serverless function to the inventory add an entry as per the below to your `inventory.yaml` file.

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
