import { ApiServiceClient } from '../protos/ApiserviceServiceClientPb';
import {
  LoginRequest,
  LoginResponse,
  LogoutRequest,
  RefreshRequest
} from '../protos/authservice_pb';
import { token } from '../utils/token';
import {
  ChangePasswordRequest,
  ChangePrivilegeRequest,
  DeleteUserRequest,
  NewUserRequest,
  SearchRequest,
  SearchResponse,
  UserResponse
} from '../protos/userservice_pb';
import {
  DeleteModelRequest,
  EditTagRequest,
  GetModelsRequest,
  GetModelsResponse,
  NewModelRequest,
  NewModelResponse
} from '../protos/modelservice_pb';
import {
  CurveFitRequest,
  CurveFitResult,
  Expression,
  RequestExpressionParameter
} from '../protos/approximationservice_pb';
import { Empty } from 'google-protobuf/google/protobuf/empty_pb';
import { Role } from '../types';
import { FitStateExpression } from '../types/stateExpression';

class ApiService {
  #client = new ApiServiceClient('http://localhost:8080');

  async #withRetry<T, R extends { setAccessToken: Function }>(
    method: T,
    request: R,
    metadata: null
  ) {
    try {
      // @ts-ignore
      return await this.#client[method](request, metadata);
    } catch (err) {
      if (err.code === 16 && err.message === 'token is expired') {
        try {
          const authRequest = new RefreshRequest();
          authRequest.setAccessToken(token.accessToken);
          authRequest.setRefreshToken(token.refreshToken);
          const response = await this.#client.refreshToken(authRequest, null);

          const {refreshToken, accessToken} = response.toObject();
          token.setRefreshToken = refreshToken;
          token.setAccessToken = accessToken;
          //@ts-ignore
          request.setAccessToken(accessToken);
          //@ts-ignore
          return await this.#client[method](request, metadata);
        } catch (err) {
          if (err.code === 16) {
            token.removeTokens();
            window.location.href = '/';
            throw err;
          }
          throw err;
        }
      } else if (err.code === 16) {
        token.removeTokens();
        window.location.href = '/';
        throw err;
      }
      throw err;
    }
  };

  //Auth Service
  Login(email: string, password: string): Promise<LoginResponse> {
    const request = new LoginRequest();
    request.setEmail(email);
    request.setPassword(password);

    return this.#client.login(request, null);
  };

  Logout(): Promise<Empty> {
    const request = new LogoutRequest();
    request.setAccessToken(token.accessToken);
    request.setRefreshToken(token.refreshToken);

    try {
      return this.#client.logout(request, null);
    } catch (e) {
      console.error(e, 'try catch works');
      throw e;
    }
  };

  // User Service
  ChangeUserPrivilege(userId: number): Promise<Empty> {
    const request = new ChangePrivilegeRequest();
    request.setUserid(userId);
    request.setNewrole(Role.ADMIN);
    request.setAccessToken(token.accessToken);

    return this.#withRetry('changeUserPrivilege', request, null);
  };

  CreateUser(name: string, email: string, password: string): Promise<UserResponse> {
    const request = new NewUserRequest();
    request.setUsername(name);
    request.setEmail(email);
    request.setPassword(password);
    request.setRole(Role.USER);
    request.setAccessToken(token.accessToken);

    return this.#withRetry('createUser', request, null);
  };

  DeleteUser(userId: number): Promise<Empty> {
    const request = new DeleteUserRequest();
    request.setId(userId);
    request.setAccessToken(token.accessToken);

    return this.#withRetry('deleteUser', request, null);
  };

  SearchForUsers(userQuery: string, keySetVal: number): Promise<SearchResponse> {
    const request = new SearchRequest();
    request.setSearchquery(userQuery);
    request.setKeysetvalue(keySetVal);
    request.setAccessToken(token.accessToken);

    return this.#withRetry('searchForUsers', request, null);
  };

  ChangePassword(userId: number, newPassword: string): Promise<Empty> {
    const request = new ChangePasswordRequest();
    request.setUserid(userId)
    request.setNewpassword(newPassword);
    request.setAccessToken(token.accessToken);

    return this.#withRetry('changePassword', request, null);
  };

  // Model Service
  AddModel(name: string, expression: string, lexexpression: string, tag: string): Promise<NewModelResponse> {
    const request = new NewModelRequest();
    request.setName(name);
    request.setExpression(expression);
    request.setLexexpression(lexexpression);
    request.setTag(tag);
    request.setAccessToken(token.accessToken);

    return this.#withRetry('addModel', request, null);
  };

  EditTag(modelId: number, tag: string): Promise<Empty> {
    const request = new EditTagRequest();
    request.setAccessToken(token.accessToken);
    request.setModelid(modelId);
    request.setNewtag(tag);

    return this.#withRetry('editTag', request, null);
  };

  DeleteModel(modelId: number): Promise<Empty> {
    const request = new DeleteModelRequest();
    request.setModelid(modelId);
    request.setAccessToken(token.accessToken);

    return this.#withRetry('deleteModel', request, null);
  };

  GetUserModels(): Promise<GetModelsResponse> {
    const request = new GetModelsRequest();
    request.setAccessToken(token.accessToken);

    return this.#withRetry('getUserModels', request, null);
  };

  // Approx Service
  FitCurves(modelsList: FitStateExpression[], xData: number[], yData: number[]): Promise<CurveFitResult> {
    // TODO: WHAT IF EMPTY IS SENT??
    const expressions: Expression[] = [];

    modelsList.forEach(({id, name, expression, lexexpression, params}) => {
      const parsedParams = params.map(param => {
        const requestParams = new RequestExpressionParameter();

        requestParams.setParamname(param.paramName);
        requestParams.setParamvalue(param.paramValue);
        requestParams.setMinbound(param.minBound);
        requestParams.setMaxbound(param.maxBound);

        return requestParams;
      });

      const newExpression = new Expression();

      newExpression.setId(id);
      newExpression.setName(name);
      newExpression.setExpression(expression);
      newExpression.setLexExpression(lexexpression);
      newExpression.setParametersList(parsedParams);

      expressions.push(newExpression);
    });

    const request = new CurveFitRequest();

    request.setExpressionsList(expressions);
    request.setXDataList(xData);
    request.setYDataList(yData);
    request.setAccessToken(token.accessToken);

    return this.#withRetry('fitCurves', request, null);
  };
}

export const apiService = new ApiService();
