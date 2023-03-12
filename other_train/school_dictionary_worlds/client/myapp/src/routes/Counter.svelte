<script>
    import {getCypherCoordinateExercise} from "../api/cypherCoordinates.js";

    import {onMount} from 'svelte';

    let matrixObj = {};
    let exerciseList = [];

    onMount(async () => {
        const res = await getCypherCoordinateExercise();
        matrixObj = res['matrix'];
        exerciseList = res['exercises'];
    });

</script>


<main>

    <ul class="pokemonList">
        {#each Object.entries(matrixObj) as [key, matrix]}
            <li>
                <label>
                    {key}
                </label>
                <label>
                    {matrix.value}
                </label>
                <label>
                    {matrix.col}
                </label>
                <label>
                    {matrix.row}
                </label>
                <label>
                    {matrix.cypher}
                </label>
            </li>
        {/each}
    </ul>

    <ul class="pokemonList">
        {#each exerciseList as exercise}
            <li>
                <label>
                    {exercise.cypher}
                </label>
                <label>
                    {exercise.answer}
                </label>
            </li>
        {/each}
    </ul>
</main>

<style>
</style>
