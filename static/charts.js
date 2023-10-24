fetch('/api/data')
    .then(response => response.json())
    .then(data => {
        // data.forEach(item => {
        //     item.date = new Date(item.date);
        // });
        const dates = data.map(item => item.date);
        const amounts = data.map(item => item.amount);

        var ctx = document.getElementById('myChart').getContext('2d');
        var myChart = new Chart(ctx, {
            type: 'bar',
            data: {
                labels: dates,
                datasets: [{
                    label: 'Debit',
                    data: amounts,
                    backgroundColor: 'rgba(75, 192, 192, 0.2)',
                    borderColor: 'rgba(75, 192, 192, 1)',
                    borderWidth: 10,
                }]
            },
            options: {
                responsive: true,
                maintainAspectRatio: true,
                title: {
                    display: true,
                    text: 'FinTracker',
                    fontSize: 18
                },
                legend: {
                    display: true,
                    position: "bottom"
                },
                scales: {
                    xAxes: [{
                        type: 'time',
                        time: {
                            tooltipFormat: "DD/MM/YY"
                        }
                    }]
                }
            }
        });
    })
    .catch(error => {
        console.error('Error:', error);
    });