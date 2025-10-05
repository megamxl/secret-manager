## Intro

This is a go program designed to work closely related to the external secret Operator but for any environment.
The goal is to write a system that can be used as a running agent or a single binary that takes 

- the file path 
- contents of the configuration file in a templating language
- And a user under which this file should be stored

## Future Ideas
- constant fetching of passwords
- spiffe spire for authentication
- Azure Key vault support 


### Plan

- Http Function that writes Files
- Templating engine that resolves static placeholders
- integrate vault secret getting 