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
                rowName: key[0],
                col: value['col'],
                value: value['value'],
                row: value['row'],
            })
        }

        matrixElements = matrixElements
    });

    let getCardValue = (row, i) => {
        const findElement = matrixElements.find(
            element => element.rowName === row && element.col === i
        )
        return findElement ? findElement.value : ''
    }

</script>


<div>
    <div class="w-5/6 m-auto">
        <div class="cypher-game m-auto">
            <div class="cypher-game__header flex">
                {#each Array(5) as _, i}
                    <CypherCard>{i > 0 ? i : ''}</CypherCard>
                {/each}
            </div>
            <div class="cypher-game__rows flex flex-wrap">
                {#each rows as row}
                    <CypherCard>{row}</CypherCard>
                    {#each Array(4) as _, i}
                        <CypherCard>{getCardValue(row, i)}</CypherCard>
                    {/each}
                {/each}
            </div>
        </div>
    </div>

    <ul>
        {#each matrixElements as matrixElement}
            <li>
                <label>
                    {matrixElement.rowName}
                    {matrixElement.row}
                    {matrixElement.value}
                </label>
            </li>

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
