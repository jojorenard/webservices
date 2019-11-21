package org.polytech;

//import io.spring.guides.gs_producing_web_service.GetBookRequest;
//import io.spring.guides.gs_producing_web_service.GetBookResponse;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.ws.server.endpoint.annotation.Endpoint;
import org.springframework.ws.server.endpoint.annotation.PayloadRoot;
import org.springframework.ws.server.endpoint.annotation.RequestPayload;
import org.springframework.ws.server.endpoint.annotation.ResponsePayload;

import org.polytech.Book;

import java.util.Map;

@Endpoint
public class BookEndpoint {
    private static final String NAMESPACE_URI = "http://spring.io/guides/gs-producing-web-service";

    private BookRepository bookRepository;

    @Autowired
    public BookEndpoint(BookRepository bookRepository){
        this.bookRepository = bookRepository;
    }

    @PayloadRoot(namespace = NAMESPACE_URI, localPart = "getBookRequest")
    @ResponsePayload
    public GetBookResponse getBook(@RequestPayload GetBookRequest request){
        GetBookResponse response = new GetBookResponse();
        response.setBook(bookRepository.findBook(request.getId()));

        return response;
    }

    @PayloadRoot(namespace = NAMESPACE_URI, localPart = "getBooksRequest")
    @ResponsePayload
    public GetBooksResponse getBooks(){
        GetBooksResponse response = new GetBooksResponse();
        Map<Integer, Book> bookMap = bookRepository.findAllBooks();
        for(Map.Entry<Integer, Book> book : bookMap.entrySet()){
            response.getBooks().add(book.getValue());
        }
        return response;
    }

}
