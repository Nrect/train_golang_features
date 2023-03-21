<script>
    import {getCypherCoordinateExercise} from "../api/cypherCoordinates.js";

    import {onMount} from 'svelte';
    import CypherCard from "../components/CypherCard.svelte";

    let matrixElements = [];
    let exerciseList = [];
    let rows = [];

    onMount(async () => {
        const res = await getCypherCoordinateExercise();
        exerciseList = res['exercises'];
        rows = res['rows'];

        for (const [key, value] of Object.entries(res['matrix'])) {
            matrixElements.push({
                row: value['cypher'][0],
                col: value['col'],
                value: value['value'],
            })
        }

        matrixElements = matrixElements
    });

</script>


<div>
    <div class="w-5/6 m-auto">
        <div class="cypher-game m-auto">
            <div class="cypher-game__header flex">
                <CypherCard></CypherCard>
                <CypherCard>1</CypherCard>
                <CypherCard>2</CypherCard>
                <CypherCard>3</CypherCard>
                <CypherCard>4</CypherCard>
            </div>
            <div class="cypher-game__rows">
                {#each rows as row}
                    <CypherCard>{row}</CypherCard>
                {/each}
            </div>
        </div>
    </div>

    <ul>
        {#each matrixElements as matrixElement}
            {#if matrixElement.row === 'А'}
                <li>
                    <label>
                        {matrixElement.row}
                        {matrixElement.col}
                        {matrixElement.value}
                    </label>
                </li>
            {/if}

        {/each}
    </ul>

    <ul>
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
</div>

<style>
    .cypher-game {
        width: 400px;
    }
</style>
