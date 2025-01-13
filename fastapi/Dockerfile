# Use the official Python image from the Docker Hub
FROM python:3.12

WORKDIR /app

RUN pip install "fastapi[standard]"

COPY app.py .

EXPOSE 8000

CMD ["uvicorn", "app:app", "--host", "0.0.0.0", "--port", "8000"]
