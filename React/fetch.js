import React, { useEffect, useState } from 'react';

function PokemonDetails() {
  const [pokemonData, setPokemonData] = useState(null);
  const [pokemonImage, setPokemonImage] = useState(null);

  useEffect(() => {
    fetch("http://127.0.0.1:3001/pokedex/1")
      .then((response) => response.json())
      .then((data) => setPokemonData(data[0]))
      .catch((error) => console.error('Error fetching data: ', error));
  }, []);
  useEffect(() => {
    fetch("http://127.0.0.1:3001/image/1")
    .then((response) => {
        if (response.ok) {
          return response.blob();
        }
        throw new Error('Network response was not ok');
      })
      .then((blob) => {
        const objectURL = URL.createObjectURL(blob);
        setPokemonImage(objectURL);
      })
      .catch((error) => console.error('Error fetching image: ', error));
  }, []);

  return (
    <div>
      <h1>Pokemon Details</h1>
      {pokemonImage ? (
        <img src={pokemonImage} alt="Pokemon" />
        ) : ( <p>Loading image...</p> )}
      {pokemonData ? (
        <div>
          <h2>{pokemonData.name}</h2>
          <p>Description: {pokemonData.description}</p>
          <p>Dex Number: {pokemonData.dexnum}</p>
          <p>Type 1: {pokemonData.type1}</p>
          <p>Type 2: {pokemonData.type2}</p>
          <p>Height: {pokemonData.height} m</p>
          <p>Weight: {pokemonData.weight} kg</p>
          <p>Base Experience: {pokemonData.base_exp}</p>
          <p>Catch Rate: {pokemonData.catch_rate}</p>
          <p>Category: {pokemonData.category}</p>
          <p>Evolve Level: {pokemonData.evolve_level}</p>
          <p>Gender Ratio: {pokemonData.gender_ratio}</p>
          <p>Stage: {pokemonData.stage}</p>
          <p>Level Speed: {pokemonData.lvl_speed}</p>
        </div>
      ) : (
        <p>Loading...</p>
      )}
    </div>
  );
}

export default PokemonDetails;
