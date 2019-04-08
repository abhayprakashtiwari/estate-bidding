- Create Estate 
```
mutation{
     createEstate(input:{
       name:"estate_1"
       description:"estate_Desc"
     }){
       id
       bids {
         amount
         bidder
       }
     }
   }
```
-Read Estate 
```
query{
     estate(id:"605394647632969758"){
       name
       description
       registeredAt
       openForBidTill
       
       bids {
         amount
         bidder
         createdAt
       }
     }
   }
```
 - Update Estate
 ```
mutation {
  updateEstate(input : {
    id : "8674665223082153551"
    name : "modified_name"
    description : "modified_desc"
  }){
    id
    name
    description
  	openForBidTill
    
  }
}
```
 - Delete Estate
 ```
mutation{
  deleteEstate(input:{
    id :"8674665223082153551"
  })
}
```

- Create Bid
```
mutation {
  createBid(input : {
    estateID : "605394647632969758"
    bidder : "bidder_6"
    amount : 1025
  }){
    id
    createdAt
  }
}
```

- Get Winner bid
```
query {
  topBid (estateID: "605394647632969758"){
    amount
    bidder   
  }
}
```