const client = require("../client_gRPC");

exports.getLogs = async (req, res) => {
  const rows = [];

  const call = client.getLogs();
  call.on("data", function (data) {
    rows.push(data);
  });
  call.on("end", function () {
    console.log("Data obtenida con exito");
    res.status(200).json({ data: rows });
  });
  call.on("error", function (e) {
    console.log("Error al obtener la data", e);
  });
};

exports.execGame = async (req, res) => {
 
  client.addLog(req.body, function (err, response) {
    res.status(200).json({ mensaje: response.message });
  });
};

