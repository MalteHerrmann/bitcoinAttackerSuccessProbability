async function updateChart() {
    let resp = await fetch('/data');
    let reply = await resp.json();
    console.log(reply)

    // Plot surface -> https://plotly.com/javascript/3d-surface-plots/
    Plotly.newPlot('chart', reply.data, reply.layout);
}

document.addEventListener('DOMContentLoaded', function() {
    // Assign updateChart function to generate-button
    document.getElementById('generate').onClick = updateChart();
})
