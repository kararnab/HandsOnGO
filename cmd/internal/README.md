# The Internal directory

## Special Importance
It’s important to point out that the directory name `internal` carries a special meaning and behavior in Go: any 
packages which live under this directory can only be imported by code inside the parent of the `internal` directory. 
In our case, this means that any packages which live in `internal` can only be imported by code inside our WebApp 
project directory. Or, looking at it the other way, this means that any packages under `internal` cannot be imported by 
code outside our project.
This is useful because it prevents other codebase from importing and relying on the (potentially un-versioned and 
unsupported) packages in our internal directory — even if the project code is publicly available somewhere like GitHub.

## Content
The internal directory will contain the ancillary non-application-specific code used in the
project. We’ll use it to hold potentially reusable code like validation helpers and the SQL
database models for the project.

