# gpu-auction

- Simple auction house for a GPU cluster
- Customers can book reservations of 1 to N GPUs, for a minimum of 1 hour
- Customers bid for the block


- What happens if a customer loses a bid? Do they lose the entire block or do they get as much as we can give them?
    - if they lose a bid, give as much we can
- How far out can someone reserve for?
    - 1 day in advance
    - when does the bid open and close?
- How do we make sure the whole GPU cluster is being used at any given time?
    - prioiritze new tasks on unused GPUs
- What happens if there’s multiple clusters, and customers who need their reservation to be on the same cluster?
    - they decide which clusters in the beginning, this will be a non-issue
- If a customer asks for multiple requests, should those be distinguished? 
    - let's make them different requests (distinguish "same")