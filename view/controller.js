/* Menu functions */

// Start
document.getElementById('start').addEventListener('click', function () {
    document.getElementById('main_menu').style.display = 'none';
    document.getElementById('name_difficulty').style.display = 'block';
});

document.getElementById('start_game').addEventListener('click', function () {
    const playerName = document.getElementById('player_name').value;
    const gameDifficulty = document.getElementById('difficulty').value;

    fetch('/createGame', {
        method: 'POST',
        body: JSON.stringify({ playerName, gameDifficulty }),
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => {
        if (response.status === 200) {
            console.log('Starting game...');
            response.json().then(data => {
                const cards = data.cards;
                console.log('Cards:', cards);
            });
        } else {
            console.error('Error on starting game: ', response.statusText);
        }
    })
    .catch(error => {
        console.error('Error on starting game: ', error);
    });

    window.location.href = `game.html`;
    const playerNameCell = "";
    const buttons = [];
    switch (gameDifficulty) {
        case 'easy':
            playerNameCell = document.querySelector('tr#first_row_easy th');
            document.getElementById('game_board_medium').style.display = 'none';
            document.getElementById('game_board_hard').style.display = 'none';
            for (let i = 0; i < 6; i++) {
                const buttonId = `card_easy_${i}`;
                const button = document.getElementById(buttonId);
                buttons.push(button);
            }
            buttons.forEach((button, index) => {
                button.setAttribute('data-card-value', cards[index]);
            });
            break;
        case 'medium':
            playerNameCell = document.querySelector('tr#first_row_medium th');
            document.getElementById('game_board_easy').style.display = 'none';
            document.getElementById('game_board_hard').style.display = 'none';
            for (let i = 0; i < 6; i++) {
                const buttonId = `card_medium_${i}`;
                const button = document.getElementById(buttonId);
                buttons.push(button);
            }
            buttons.forEach((button, index) => {
                button.setAttribute('data-card-value', cards[index]);
            });
        case 'hard':
            playerNameCell = document.querySelector('tr#first_row_hard th');
            document.getElementById('game_board_easy').style.display = 'none';
            document.getElementById('game_board_medium').style.display = 'none';
            for (let i = 0; i < 6; i++) {
                const buttonId = `card_hard__${i}`;
                const button = document.getElementById(buttonId);
                buttons.push(button);
            }
            buttons.forEach((button, index) => {
                button.setAttribute('data-card-value', cards[index]);
            });
            break;
        default:
            break;
    }
    playerNameCell.textContent = playerName;

});

// Ranking
document.getElementById('ranking').addEventListener('click', function () {
    fetch('/getRanking', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => response.json())
        .then(data => {
            const gameRankingDiv = document.getElementById('game_ranking');
            gameRankingDiv.innerHTML = JSON.stringify(data, null, 2);
        })
        .catch(error => {
            console.error('Error obtaining ranking:', error);
        });    
});

// Exit
document.getElementById('exit').addEventListener('click', function () {
    if (confirm('Are you sure you want to exit?')) {
        window.close();
    }
});

