let PLAYER_NAME;

const getCardImage = (value) => {
    switch (value) {
        case 0:
            return "clover-10.jpg"
        case 1:
            return "clover-j.jpg"
        case 2:
            return "clover-q.jpg"
        case 3:
            return "clover-k.jpg"
        case 4:
            return "clover-as.jpg"
        case 5:
            return "diamond-10.jpg"
        case 6:
            return "diamond-j.jpg"
        case 7:
            return "diamond-q.jpg"
        case 8:
            return "diamond-k.jpg"
        case 9:
            return "diamond-as.jpg"
        case 10:
            return "spades-10.jpg"
        case 11:
            return "spades-j.jpg"
        case 12:
            return "spades-q.jpg"
        case 13:
            return "spades-k.jpg"
        case 14:
            return "spade-as.jpg"
        case 15:
            return "heart-10.jpg"
        case 16:
            return "heart-j.jpg"
        case 17:
            return "heart-q.jpg"
        case 18:
            return "heart-k.jpg"
        case 19:
            return "heart-as.jpg"
        default:
            return "card-back.jpg"
    }
};
const getRowCellsByDifficulty = (value) => {
    switch (value) {
        case 0:
            return 3;
        case 1:
            return 4;
        default:
            return 4;
    }
}

const getData = async (url, body) => {
    try {
        const response = await fetch(url, {
            method: 'POST',
            body: body,
            headers: {
                'Content-Type': 'application/json'
            }
        });
        return response.json();
    } catch (error) {
        console.error(error);
    }
}

const setGameBoard = (cards, difficulty) => {
    const board = $('#tableBody');
    const cells = cards.length;
    board.html('');

    for (let i = 0; i < cells; i++) {
        if (i === 0 || i % getRowCellsByDifficulty(difficulty) === 0) {
            const newRow = document.createElement('tr');
            board.append(newRow);
        }

        const newCell = document.createElement('td');
        const card = document.createElement('button');
        const backCard = document.createElement('img');

        backCard.id = "card_img_" + i;
        backCard.style.maxWidth = '250px';
        backCard.classList.add('img-fluid');
        console.log(cards[i].disable);
        if (!cards[i].disable)
            backCard.src = "/icons/card-back.jpg"
        else
            backCard.src = '/icons/card-blank.jpg'
        card.id = "card_" + i;
        card.classList.add('btn', 'btn-outline-primary');
        card.setAttribute('card-id', i)
        card.appendChild(backCard);
        card.addEventListener('click', chooseCard)

        newCell.appendChild(card)
        board.children().last().append(newCell)
    }
}

const startGame = async (gameDifficulty) => {
    const url = 'http://localhost:8080/createGame';
    const data = await getData(url, JSON.stringify({player_name: PLAYER_NAME, difficulty: parseInt(gameDifficulty)}));

    console.log('Starting game...');
    console.log('Game difficulty: ', gameDifficulty);
    const cards = data.cards;
    console.log(cards)
    const buttons = [];

    setGameBoard(cards, gameDifficulty);
    updateCards(cards);
    $('#endGameButton').css('display', 'block')
}

document.getElementById('start_game').addEventListener('click', function () {
    PLAYER_NAME = document.getElementById('player_name').value;
    const gameDifficulty = document.getElementById('difficulty').value;
    const tableDifficulties = document.getElementsByClassName('playerName');
    Array.from(tableDifficulties).forEach(element => {
        element.textContent = PLAYER_NAME;
    });
    startGame(parseInt(gameDifficulty));
});

// Click on card
function chooseCard(event) {
    console.log("Choosing a Card");
    const buttonId = event.target.id;
    const card_id = buttonId.split('_')[2];
    const url = 'http://localhost:8080/chooseCard';
    console.log("Selected card: ", card_id);
    getChosenCard(url, parseInt(card_id));
}

const getChosenCard = async (url, card_id) => {
    const body = JSON.stringify({player_name: PLAYER_NAME, card_choice: card_id});
    const data = await getData(url, body);

    if (data.success) {
        $('#card_img_' + card_id).attr('src', "/icons/" + getCardImage(data.cards[card_id].value));
        updateCards(data.cards);
    }
}

const updateCards = (cards) => {
    console.log("Update Cards")
    let disabledCards = [];

    cards.forEach((card, index) => {
        if (card.visible && !card.disable) {
            const visibleCard = document.getElementById("card_img_" + index);
            visibleCard.src = "/icons/" + getCardImage(card.value);
        }
        if (card.disable) {
            disabledCards.push($('#card_img_' + index));
        }
    });

    setTimeout(() => {
        let opacity = 1;
        const opacityInterval = setInterval(() => {
            opacity -= 0.01;
            disabledCards.forEach((image) => {
                image.css('opacity', opacity)
            });
            if (opacity <= 0) {
                clearInterval(opacityInterval);
                disabledCards.forEach((image) => {
                    image.attr("src", "/icons/card-blank.jpg");
                    image.parent().prop('disabled', true)
                });
            }
        }, 20);
    }, 1000);
}

const handleEnd = async () => {
    const url = 'http://localhost:8080/end'
    const body = JSON.stringify({player_name: PLAYER_NAME});
    const data = await getData(url, body);

    console.log("points: ", data.points);
    $('#board').prepend(alertMessage('success', 'Game ended!',
        'You can start another game. Points: ' + data.points))

}
const alertMessage = (level, title, message) => (
    '<div class="alert alert-' + level + ' alert-dismissible fade show" role="alert">' +
    '   <strong>' + title + '</strong> ' + message +
    '   <button type="button" class="close" data-dismiss="alert" aria-label="Close">' +
    '       <span aria-hidden="true">&times;</span>' +
    '   </button>' +
    '</div>')
const updateScore = async () => {
    const url = 'http://localhost:8080/getRanking'
    const body = JSON.stringify({player_name: PLAYER_NAME});
    const data = await getData(url, body);

    const playerWithPoints = data.players.find(player =>
        player.player_name === PLAYER_NAME && player.points !== undefined
    );
}