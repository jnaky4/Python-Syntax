import fetch from "node-fetch";
globalThis.fetch = fetch

//base example
const fetchData = () => {
    fetch('https://api.github.com').then(resp => { //fetch: async function, returns a promise, we do a .then to consume the promise on the result of fetch
        resp.json().then(data => {  //resp returns a Raw response, if you want json, you call json to return it as a json, which also returns a promise
            console.log(data);
        });
    });
};
fetchData();

//modern away, async await
const fetchData2 = async () => { //syntax to remove .then calls
    const resp = await fetch('https://api.github.com');
    const data = await resp.json();
    console.log(data);
}
fetchData2();
