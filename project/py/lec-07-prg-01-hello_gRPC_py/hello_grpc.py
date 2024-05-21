def my_func(input_number):
    return input_number * input_number

# 클래스 화일 생성
# -> python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. filename.proto