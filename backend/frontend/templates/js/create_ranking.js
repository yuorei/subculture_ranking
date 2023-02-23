const btn = document.querySelector('.btn');
const pass = window.location.pathname
const words = pass.split('/');
let url = '/ranking/' + words[2]


const postFetch = () => {
    const data = {
        Title: document.forms['form1'].elements['name'].value,
        UserId: Number(words[2]),
        Rank: Number(document.forms['form1'].elements['rank'].value),
        PosterComment: document.forms['form1'].elements['comment'].value,
    };

    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
        // error処理
        .then((response) => {
            if (!response.ok) {
                throw new Error(`リクエストに失敗 : $(response.status)`)
            }
            return response.json()
        })
        .then((data) => {
            console.log('Success:', data);
            alert("成功");
            let form = document.getElementById("form");
            form.reset();
        })
        .catch((error) => {
            console.error('Error:', error);
            alert("タイトルを見つけられませんでした\n別のものを入力してください");
        });
};

btn.addEventListener('click', postFetch, false);