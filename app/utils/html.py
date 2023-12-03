from app.models.ServerMetrics import ServerMetrics


def format_timestamp(timestamp):
    return timestamp[11:19] + " (UTC)"


def generate_server_status_html(metrics: ServerMetrics) -> str:
    return f"""
    <!DOCTYPE html>
<html lang=en>
<head>
    <meta charset=UTF-8>
    <meta name=viewport content=width=device-width, initial-scale=1.0>
    <style>
        @import url('https://fonts.googleapis.com/css?family=Open+Sans&display=swap');
        html {{
            position: relative;
            overflow-x: hidden !important;
        }}
        * {{
            box-sizing: border-box;
        }}
        body {{
            font-family: 'Open Sans', sans-serif;
            color: #324e63;
        }}
        a {{
            color: inherit;
            text-decoration: inherit;
        }}
        .wrapper {{
            width: 100%;
            width: 100%;
            height: auto;
            min-height: 90vh;
            padding: 50px 20px;
            padding-top: 100px;
            display: flex;
        }}
        .instance-card {{
            width: 100%;
            min-height: 380px;
            margin: auto;
            box-shadow: 12px 12px 2px 1px rgba(13, 28, 39, 0.4);
            background: #fff;
            border-radius: 15px;
            border-width: 1px;
            max-width: 500px;
            position: relative;
            border: thin groove #9c83ff;
        }}
        .instance-card__cnt {{
            margin-top: 35px;
            text-align: center;
            padding: 0 20px;
            padding-bottom: 40px;
            transition: all .3s;
        }}
        .instance-card__name {{
            font-weight: 700;
            font-size: 24px;
            color: #6944ff;
            margin-bottom: 15px;
        }}
        .instance-card-inf__item {{
            padding: 10px 35px;
            min-width: 150px;
        }}
        .instance-card-inf__title {{
            font-weight: 700;
            font-size: 27px;
            color: #324e63;
        }}
        .instance-card-inf__txt {{
            font-weight: 500;
            margin-top: 7px;
        }}
        .secondary{{color: #9c83ff;}}
    </style>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.2.1/css/all.min.css" integrity="sha512-MV7K8+y+gLIBoVD59lQIYicR65iaqukzvf/nwasF0nqhPay5w/9lJmVM2hMDcnK1OnMGCdVK+iQrJ7lzPJQd1w==" crossorigin="anonymous" referrerpolicy="no-referrer" />
    <title>😈️ pointless-status 😈</title>
</head>
<body>
    <div class=wrapper>
        <div class=instance-card>
            <div class=instance-card__cnt>
                <div class=instance-card__name>😈️ Server is running! 😈️</div>
                <div class=instance-card-inf>
                    <div class=instance-card-inf__item>
                        <div class=instance-card-inf__txt>Client</div>
                        <div class=instance-card-inf__title>{metrics.host}</div>
                    </div>
                    <div class=instance-card-inf__item>
                        <div class=instance-card-inf__txt>Time</div>
                        <div class=instance-card-inf__title>{format_timestamp(metrics.timestamp)}</div>
                    </div>
                    <div class=instance-card-inf__item>
                        <div class=instance-card-inf__txt>CPU Usage</div>
                        <div class=instance-card-inf__title>{metrics.cpu_usage}</div>
                    </div>
                    <div class=instance-card-inf__item>
                        <div class=instance-card-inf__txt>Memory usage</div>
                        <div class=instance-card-inf__title>{metrics.memory_usage}</div>
                    </div>
                    <div class=instance-card-inf__item>
                        <div class="instance-card-inf__title secondary"><a href="https://github.com/linomp" target="_blank"><i class="fab fa-github"></i></a></div>
                        <div class="instance-card-inf__txt secondary"><a href="/docs" target="_blank">/docs</a></div>
                    </div>
                </div>
            </div>
        </div>
</body>
</html>
    """
