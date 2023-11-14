document.getElementById('start_game').addEventListener('click', function () {
    const playerName = document.getElementById('player_name').value;
    const gameDifficulty = document.getElementById('difficulty').value;

    fetch('http://localhost:8080/createGame', {
        method: 'POST',
        body: JSON.stringify({ player_name: playerName , difficulty:parseInt(gameDifficulty) }),
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => {
        if (response.status === 201) {
            console.log('Starting game...');
            response.json().then(data => {
                var cards = data.cards;
                console.log('Cards:', cards);
                let playerNameCell;
                const buttons = [];
                switch (gameDifficulty) {
                    case 0:
                        playerNameCell = document.querySelector('th#easy_player_name');
                        document.getElementById('game_board_medium').style.display = 'none';
                        document.getElementById('game_board_hard').style.display = 'none';
                        for (let i = 0; i < 6; i++) {
                            const buttonId = `card_easy_${i}`;
                            const button = document.getElementById(buttonId);
                            buttons.push(button);
                        }
                        buttons.forEach((button, index) => {
                            button.setAttribute('data-card-value', cards[index].value);
                        });
                        break;
                    case 1:
                        playerNameCell = document.querySelector('th#medium_player_name');
                        document.getElementById('game_board_easy').style.display = 'none';
                        document.getElementById('game_board_hard').style.display = 'none';
                        for (let i = 0; i < 6; i++) {
                            const buttonId = `card_medium_${i}`;
                            const button = document.getElementById(buttonId);
                            buttons.push(button);
                        }
                        buttons.forEach((button, index) => {
                            button.setAttribute('data-card-value', cards[index].value);
                        });
                        break;
                    case 2:
                        playerNameCell = document.querySelector('th#hard_player_name');
                        document.getElementById('game_board_easy').style.display = 'none';
                        document.getElementById('game_board_medium').style.display = 'none';
                        for (let i = 0; i < 6; i++) {
                            const buttonId = `card_hard__${i}`;
                            const button = document.getElementById(buttonId);
                            buttons.push(button);
                        }
                        buttons.forEach((button, index) => {
                            button.setAttribute('data-card-value', cards[index].value);
                        });
                        break;
                    default:
                        break;
                }
                playerNameCell.textContent = data.player_name;
            });
        } else {
            console.error('Error on starting game: ', response.statusText);
        }
    })
    .catch(error => {
        console.error('Error on starting game: ', error);
    });
});

// Click on card
function handleClick(event) {
    const buttonId = event.target.id;
    const cardValue = event.target.getAttribute('data-card-value');

    fetch('http://localhost:8080/chooseCard', {
        method: 'POST',
        body: JSON.stringify({ player_name, card_choice }),
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => {
        if (response.status === 200) {
            response.json().then(data => {
                if (data.Success === true) {
                    // Mostrar valor de la carta
                    //document.getElementById(buttonId).style.display = 'none';
                }
                console.log('Cards:', data.cards);
            });
        } else {
            console.error('Error on starting game: ', response.statusText);
        }
    })
    .catch(error => {
        console.error('Error on starting game: ', error);
    });


}
document.querySelectorAll('button').forEach(button => {
    button.addEventListener('click', handleClick);
});