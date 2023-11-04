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
        } else {
            console.error('Error on starting game: ', response.statusText);
        }
    })
    .catch(error => {
        console.error('Error on starting game: ', error);
    });
    window.location.href = `game.html`;
    const playerNameCell = document.querySelector('table th');
    playerNameCell.textContent = playerName;
});

// Ranking
document.getElementById('ranking').addEventListener('click', function () {
    // Lógica para mostrar el ranking
    alert('Showing ranking...');
});

// Exit
document.getElementById('exit').addEventListener('click', function () {
    if (confirm('Are you sure you want to exit?')) {
        window.close();
    }
});

