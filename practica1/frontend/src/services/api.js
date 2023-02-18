const URL_API = `http://192.168.0.10:4000`;
// const URL_API = `${process.env.URL_API}`;



const suma = async (data) => {
  const response = await fetch(`${URL_API}/suma`, {
    method: 'POST',
    body: JSON.stringify(data),
    headers: {
      'Accept':'application/json',
      'Content-Type': 'application/json',
    }    
  });
  const responseJson = await response.json();
  // console.log(responseJson);
  return responseJson.result
};


const resta = async (data) => {
  const response = await fetch(`${URL_API}/resta`, {
    method: 'POST',
    body: JSON.stringify(data),
    headers: {
      'Accept':'application/json',
      'Content-Type': 'application/json',
    }    
  });
  const responseJson = await response.json();
  // console.log(responseJson);
  return responseJson.result
};

const multiplicacion = async (data) => {
  const response = await fetch(`${URL_API}/multiplicacion`, {
    method: 'POST',
    body: JSON.stringify(data),
    headers: {
      'Accept':'application/json',
      'Content-Type': 'application/json',
    }    
  });
  const responseJson = await response.json();
  // console.log(responseJson);
  return responseJson.result
};

const division = async (data) => {
  const response = await fetch(`${URL_API}/division`, {
    method: 'POST',
    body: JSON.stringify(data),
    headers: {
      'Accept':'application/json',
      'Content-Type': 'application/json',
    }    
  });
  const responseJson = await response.json();
  // console.log(responseJson);
  return responseJson.result
};


export const historial = async () => {
  const response = await fetch(`${URL_API}/historial`,{
    method: 'GET',
    headers: {
      'Accept':'application/json',
      'Content-Type': 'application/json',
    }
  });
  const responseJson = await response.json();
  return responseJson;
};


export const operar = (value1, value2, operador) => {
  const data = {value1: parseInt(value1), value2: parseInt(value2)};
  // console.log(data);
  switch (operador) {
    case "+":
      return suma(data);
      
    case "-":
      return resta(data);  

    case "*":
      return multiplicacion(data);  

    case "/":
      return division(data);
  }
};
