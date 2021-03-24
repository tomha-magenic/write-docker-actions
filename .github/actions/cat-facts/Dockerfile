FROM python:3

COPY requirements.txt ./

RUN pip install --nocache-dir -r requirements.txt

COPY . .

CMD ["python", "/src/main.py"]