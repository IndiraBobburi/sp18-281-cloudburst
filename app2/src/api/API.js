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
        credentials:'include',
        body: JSON.stringify(payload)
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
        credentials:'include',
        body: JSON.stringify(payload)
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
            return error;
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
            return error;
        });

export const addToCart = (payload) =>
    fetch(`${api_server}/addToCart`, {
        method: 'POST',
        headers: {
            ...headers,
            'Content-Type': 'application/json'
        },
        credentials:'include',
        body: JSON.stringify(payload)
    }).then(res => {
        //localStorage.setItem("cart",JSON.stringify(payload));
        return res.json();
    })
        .catch(error => {
            console.log("This is error");
            return error;
        });


export const getCart = (payload) =>
    fetch(`${api_server}/viewCart?userId=`+payload, {
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
           // return res;
        });

export const order = (payload) =>
    fetch(`${api_server}/order`, {
        method: 'POST',
        headers: {
            ...headers,
            'Content-Type': 'application/json'
        },
        credentials:'include',
        body: JSON.stringify(payload)
    }).then(res => {
        //localStorage.setItem("cart",JSON.stringify(payload));
        return res.json();
    })
        .catch(error => {
            console.log("This is error");
            // localStorage.setItem("cart",JSON.stringify(payload));
            // return error;
            return error;
        });
export const getAllOrders = (payload) =>
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