import Api from "../Api";

export const getCypherCoordinateExercise = async () => {
    try {
        const response = await Api.get("/cypher_coordinates");
        return response;
    } catch (error) {
        console.error(error);
    }
};