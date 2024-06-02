import sqlite from 'sqlite3';
const { verbose } = sqlite;

const sqlite3 = verbose();

const DATABASE_PATH = '../../../database/speedtest-1.db'

const db = new sqlite3.Database(DATABASE_PATH, (err) => {
    if (err) {
        console.error(err.message);
    }
    console.log('Connected to the speedtest database.');
})

const search = () => {
  const allRows = [];
  db.all("SELECT * FROM speedtest", (err, rows) => {
    if (err) {
      throw err;
    }
    // allRows = rows;
    rows.forEach((row) => {
      console.log(row);
    });
  });

  return allRows;
}

search();

db.close();
