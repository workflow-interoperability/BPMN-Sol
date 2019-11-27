pragma solidity ^0.4.24;

contract bpmn_v1 {
    uint public status = 0;  //identify which step is now
    //0 for start, 1 for finish order and wait for seller to recieve
    //2 for recieved and wait for the goods to be delievered
    address sellerAddr = 0xb820b9f01be068e273ac75e530f1f55c70cc648d;
    uint public currentPrice = 1000;
    
    struct order{
        uint id;
        uint amount;
        uint price;
        uint totalPrice;
        bool recieved;
        address addr;  //address of the buyer
    }
    
    event ordersent(uint id, address buyer);
    event orderrecieved(uint id);

    order[] public orders;  //a set of orders
    
    function setPrice(uint newPrice) {  //set the price, call by the seller
        require(msg.sender == sellerAddr); 
        currentPrice = newPrice;
    }
    
    function orderGoods(uint id_, uint amount_) {  //order goods, call by the buyers
        require(msg.sender != sellerAddr);
        require(status == 0);
        orders.push(order({
            id: id_,
            amount: amount_,
            price: currentPrice,
            totalPrice: amount_ * currentPrice,
            recieved: false,
            addr: msg.sender
        }));
        status = 1;  //order is sent by the buyer, waiting for the seller to recieve
        ordersent(id_, msg.sender);
    }
    
    function recieveOrders(uint id_) {  //recieve order, call by the seller
        require(msg.sender == sellerAddr);
        require(status == 1);
        //confirm
        uint i;
        for (i = 0; i < orders.length; i++) {
            if (orders[i].id == id_ && orders[i].recieved == false) {
                orders[i].recieved = true;
                status = 2;  //order recieved
                orderrecieved(id_);
            }
        }
    }
    
    function recieveGoods(uint id_) {  //confirm recieving the goods, call by the buyer who order it
        require(status == 2);
        //confirm
        uint i;
        for (i = 0; i < orders.length; i++) {
            if (orders[i].id == id_ && orders[i].recieved == true && orders[i].addr == msg.sender) { 
                status = 0;  //back to the starting status
            }
        }
    }
    
}

