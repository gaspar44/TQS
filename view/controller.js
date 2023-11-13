/* Menu functions */

// Start
document.getElementById('start').addEventListener('click', function () {
    document.getElementById('main_menu').style.display = 'none';
    document.getElementById('name_difficulty').style.display = 'block';
});

document.getElementById('start_game').addEventListener('click', function () {
    const playerName = document.getElementById('player_name').value;
    const gameDifficulty = document.getElementById('difficulty').value;

    window.location.href = `game.html`;
    var playerNameCell = "";
    const buttons = [];

    // Setting board
    switch (gameDifficulty) {
        case 'easy':
            playerNameCell = document.querySelector('th#easy_player_name');
            document.getElementById('game_board_medium').style.display = 'none';
            document.getElementById('game_board_hard').style.display = 'none';
            break;
        case 'medium':
            playerNameCell = document.querySelector('th#medium_player_name');
            document.getElementById('game_board_easy').style.display = 'none';
            document.getElementById('game_board_hard').style.display = 'none';
            break;
        case 'hard':
            playerNameCell = document.querySelector('th#hard_player_name');
            document.getElementById('game_board_easy').style.display = 'none';
            document.getElementById('game_board_medium').style.display = 'none';
            break;
        default:
            break;
    }
    playerNameCell.textContent = playerName;

    fetch('http://localhost:8080/createGame', {
        method: 'POST',
        body: JSON.stringify({ player_name: playerName , difficulty:parseInt(gameDifficulty) }),
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => {
        if (response.status === 201) {
            response.json().then(data => {
                console.log('Data: ', data)
                const cards = data.cards;   
                
                // Setting cards
                switch (gameDifficulty) {
                    case 'easy':
                        for (let i = 0; i < length(cards); i++) {
                            const buttonId = `card_easy_${i}`;
                            const button = document.getElementById(buttonId);
                            buttons.push(button);
                        }                        
                        break;
                    case 'medium':
                        for (let i = 0; i < length(cards); i++) {
                            const buttonId = `card_medium_${i}`;
                            const button = document.getElementById(buttonId);
                            buttons.push(button);
                        }
                        break;
                    case 'hard':
                        for (let i = 0; i < length(cards); i++) {
                            const buttonId = `card_hard__${i}`;
                            const button = document.getElementById(buttonId);
                            buttons.push(button);
                        }
                        break;
                    default:
                        break;
                }
                // Assign card values
                buttons.forEach((button, index) => {
                    button.setAttribute('data-card-value', cards[index].Value);
                });
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
                    const buttonId = event.target.id;
                    const cardValue = event.target.getAttribute('data-card-value');
                    document.getElementById(buttonId).textContent = cardValue;
                } else {
                    const board = document.getElementById("board"); 
                    const buttons = board.querySelectorAll("button");
                    
                    buttons.forEach((button) => {
                        const isButtonVisible = window.getComputedStyle(button).display !== "none";   

                        if (isButtonVisible) {
                            const dataCardValue = button.getAttribute("data-card-value");
                            if (dataCardValue === valorComparar) {
                                document.getElementById(buttonId).textContent = buttonId.value;
                            }
                        }
                    })
                }
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

// Ranking
document.getElementById('ranking').addEventListener('click', function () {
    fetch('http://localhost:8080/getRanking', {
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

