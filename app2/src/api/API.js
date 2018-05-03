const api_server ='http://localhost:8080'

const headers = {
    'Accept': 'application/json'
};

export const signin = (payload) =>
    fetch(`${api_server}/signin`, {
        method: 'POST',
        headers: {
            ...headers,
            'Content-Type': 'application/json'
        },
        credentials:'include'
    }).then(res => {
        return res.json();
    })
        .catch(error => {
            console.log("This is error");
            return error;
        });

export const signup = (payload) =>
    fetch(`${api_server}/signup`, {
        method: 'POST',
        headers: {
            ...headers,
            'Content-Type': 'application/json'
        },
        credentials:'include'
    }).then(res => {
        return res.json();
    })
        .catch(error => {
            console.log("This is error");
            return error;
        });

export const getRestaurants = (payload) =>
    fetch(`${api_server}/getRestaurants?pincode=`+payload, {
        method: 'GET',
        headers: {
            ...headers,
            'Content-Type': 'application/json'
        },
        credentials:'include'
    }).then(res => {
        return res.json();
    })
        .catch(error => {
            console.log("This is error");
           // return error;
            var res = {
                "restaurantlist" : [
                    {
                        "id": 1,
                        "name": "mcd",
                        "address": "xyz",
                        "phone": "320-234-2384"
                    },
                    {
                        "id": 2,
                        "name": "burgerking",
                        "address": "abc",
                        "phone": "320-234-3456"
                    }
                ]
            }
            return res;
        });

export const getMenu = () =>
    fetch(`${api_server}/getMenu`, {
        method: 'GET',
        headers: {
            ...headers,
            'Content-Type': 'application/json'
        },
        credentials:'include'
    }).then(res => {
        return res.json();
    })
        .catch(error => {
            console.log("This is error");
            // return error;
            var res = {
                "menu" : [
                    {
                        "id": 1,
                        "name": "veggie burger",
                        "price": 14.0,
                        "description": "Homemade delicious burger with roasted onions, peppers"
                    },
                    {
                        "id": 2,
                        "name": "cheese burger",
                        "price": 16.0,
                        "description": "Burger with finely chopped cheese"
                    },
                    {
                        "id": 3,
                        "name": "bacon burger",
                        "price": 16.0,
                        "description": "bacon Burger"
                    },
                    {
                        "id": 2,
                        "name": "steak burger",
                        "price": 18.0,
                        "description": "steak Burger with finely chopped cheese"
                    }
                ]
            }
            return res;
        });

export const getOrders = (payload) =>
    fetch(`${api_server}/orders`, {
        method: 'GET',
        headers: {
            ...headers,
            'Content-Type': 'application/json'
        },
        credentials:'include'
    }).then(res => {
        return res.json();
    })
        .catch(error => {
            console.log("This is error");
            return error;
        });