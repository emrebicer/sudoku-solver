// Generate the Sudoku grid dynamically
const sudokuGrid = document.getElementById('sudokuGrid');
for (let i = 0; i < 81; i++) {
	const input = document.createElement('input');
	input.type = 'text';
	input.className = 'sudoku-input';
	input.maxLength = 1;
	input.addEventListener('input', function(_e) {
		// Ensure only digits 1-9 are entered
		this.value = this.value.replace(/[^1-9]/g, '');
		this.style.color = 'white';
	});

	let borderProperty = '2px solid #D24317';

	// Add thicker borders
	if (i < 9) {
		input.style.borderTop = borderProperty;
	}
	if (i > 71) {
		input.style.borderBottom = borderProperty;
	}
	if ((i + 1) % 9 === 0) {
		input.style.borderRight = borderProperty;
	}
	if ((i + 1) % 9 === 1) {
		input.style.borderLeft = borderProperty;
	}
	if (((i + 1) % 3 === 0) && ((i + 1) % 9 !== 0)) {
		input.style.borderRight = borderProperty;
	}
	if (i >= 18 && i < 27 || i >= 45 && i < 54) {
		input.style.borderBottom = borderProperty;
	}

	sudokuGrid.appendChild(input);
}

function generateSudokuBoardString() {
	const sudokuGrid = document.getElementById('sudokuGrid');
	const inputs = sudokuGrid.querySelectorAll('.sudoku-input');
	let sudokuString = '';

	inputs.forEach((input) => {
		const value = input.value === '' ? '0' : input.value;
		sudokuString += value;
	});

	return sudokuString;
}

function fillSudokuGridFromSolutionString(solutionString) {

	if (solutionString.length !== 81) {
		console.error("The input string must be 81 characters long: ", solutionString);
		alert("failed to fill the sudoku grid")
		return;
	}

	let currentBoardString = generateSudokuBoardString();

	const sudokuGrid = document.getElementById('sudokuGrid');
	const inputs = sudokuGrid.querySelectorAll('.sudoku-input');

	solutionString.split('').forEach((char, index) => {
		if (currentBoardString[index] == "0") {
			inputs[index].style.color = '#00A66E';
			inputs[index].value = char;
		}
	});
}


function clearSudokuBoard() {
	const sudokuGrid = document.getElementById('sudokuGrid');
	const inputs = sudokuGrid.querySelectorAll('.sudoku-input');

	inputs.forEach(input => {
		input.style.color = 'white';
		input.value = '';
	});
}

async function solveSudoku() {
	// Disable all buttons
	const allBtn = document.querySelectorAll('button')
	allBtn.forEach((button) => {
		button.disabled = true;
	});

	// Generate the string representation of the Sudoku board
	let sudokuString = generateSudokuBoardString();

	try {
		const response = await fetch(`/solve?board=${sudokuString}`, { signal: AbortSignal.timeout(3000) });

		if (!response.ok) {
			throw new Error(`Error: ${response.statusText}`);
		}

		const solutionString = await response.text();
		fillSudokuGridFromSolutionString(solutionString)

	} catch (error) {
		// Show an alert with the error message
		alert(error);
	}

	// Enable all buttons
	allBtn.forEach((button) => {
		button.disabled = false;
	});
}
