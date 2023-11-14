/* Menu functions */

// Start
document.getElementById('start').addEventListener('click', function () {
    window.location.href = `game.html`;
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