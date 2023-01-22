const btn = document.querySelector('.btn');
const pass=window.location.pathname
const words = pass.split('/');
let url = '/ranking/'+words[2]


const postFetch = () => {
    const data = {
        Title:  document.forms['form1'].elements['name'].value,
        UserId: Number(words[2]),
        Rank:   Number(document.forms['form1'].elements['rank'].value),
    };

    fetch( url, {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
    })
    // error処理
    .then((response) => response.json())
    .then((data) => {
        console.log('Success:', data);
    })
    .catch((error) => {
        console.error('Error:', error);
    });
    };

btn.addEventListener('click', postFetch, false);