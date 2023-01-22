const btn = document.querySelector('.btn');
const url = '/users';

const postFetch = () => {
    const data = {
        Name:  document.forms['form1'].elements['name'].value,
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