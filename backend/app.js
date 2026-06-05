const express = require("express");
const cors = require("cors");

const app = express();
const port = 3000;

app.use(cors());
app.use(express.json());

const path = require("path");

app.use(express.static(path.join(__dirname, "../frontend")));

app.use("/products", routesProduits);

app.listen(port, () => {
  console.log("Serveur lancé sur http://localhost:" + port);
});