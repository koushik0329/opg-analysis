# OPG Analysis CLI

A command-line tool for performing stock analysis and making trade decisions based on stock data, news articles, and customizable parameters.

---

## **Features**
- Load stock data from CSV files.
- Fetch news articles for stock tickers.
- Calculate trade positions based on input parameters.
- Generate an output JSON file with analysis results.

---

## **Installation**

Clone the Repository
   ```bash
   git clone https://github.com/koushik0329/opg-analysis.git
   cd opg-analysis
   ```
Build the Executable

    go build -o opg .
Run the Executable To see available command options:

    ./opg
Environment Variables
Before running the application, set the required environment variables:

SEEKING_ALPHA_URL: The API URL for fetching news articles.

SEEKING_ALPHA_API_KEY: The API key for accessing the Seeking Alpha API.

generate api key using rapid api
Example:

    export SEEKING_ALPHA_URL=https://seeking-alpha.p.rapidapi.com
    export SEEKING_ALPHA_API_KEY=your_api_key
Usage
Command Options

    ./opg -i <input-path> -b <account-balance> [optional-flags]


---

Flag	Description	Example

    ./opg

| Flag | Description                                        | Example              |
|------|----------------------------------------------------|----------------------|
| `-i` | Path to the input CSV file (required).             | `-i ./opg.csv`       |
| `-b` | Account balance for trade calculations (required). | `-b 1000`            |
| `-o` | Path to output JSON file (default: `./opg.json`).  | `-o ./results.json`  |
| `-m` | Minimum gap value to consider (default: `0.1`).    | `-m 0.2`             |
| `-l` | Loss tolerance percentage (default: `0.02`).       | `-l 0.05`            |
| `-p` | Profit percentage of the gap to take (default: `0.8`). | `-p 0.7`         |

Examples

Run Analysis with Default Parameters

    ./opg -i ./opg.csv -b 1000
    
    ./opg -i ./opg.csv -b 1000 -m 0.2


---

### Workflow
Load Stock Data: Reads data from a CSV file with stock tickers, gaps, and opening prices.

Fetch News Articles: Uses the Seeking Alpha API to fetch relevant news articles for each stock.

Calculate Trade Positions: Uses input parameters like account balance, loss tolerance, profit percentage, and minimum gap to compute trade positions.

Save Results: Outputs the analysis results to a JSON file.

---
### Example Input File
opg.csv

---
### Output
After running the analysis, the output will be saved as a JSON file (default: opg.json).

---

## **Source**
https://youtu.be/0hChKDYOKd8?si=uLsWvqT_TqPI0LRH
