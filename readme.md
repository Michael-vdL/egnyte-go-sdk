# Egnyte Go SDK

Reusable Go Client for Egnyte's API.

# Roadmap

Initial Release is targetted at Events API.
Working into a corporate project to collect Egnyte events and log them in Elastic Search.

After initial Events are defined, will work through remaining endpoints as needed.

# Requirements

- CMD Line Usage
- Reusable in consumer serivce
- Does not rely on libraries outside of Go STD.

# Notes

- I want to refactor the code so that a single client type can make all the necessary calls, however, code is functional for current usecases and I am short on time.
- When implementing SDK, follow Egnyte API best practices
  - Cache user queries and maintain reference to them outside of API.
  - Call events API once every 5 minutes if possible.
  - Read https://developers.egnyte.com/docs

# Tags

- 0.1.3 User Query Paramets
